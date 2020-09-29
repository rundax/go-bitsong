package cli

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bitsongofficial/go-bitsong/x/mpeg21/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"
	"io/ioutil"
	"strings"
)

func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	contentTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	contentTxCmd.AddCommand(flags.PostCommands(
		GetCmdStoreMCO(cdc),
	)...)

	return contentTxCmd
}

func GetCmdStoreMCO(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "store-mco",
		Short: "Store a new Smart Media Contract (MPEG21 MCO)",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Store a new Smart Media Contract (MPEG21 MCO).
Example:
$ %s tx mpeg21 store-mco [contract.json] --from <creator>`,
				version.ClientName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			contractInfo, err := ioutil.ReadFile(args[0])
			if err != nil {
				return err
			}

			contractInfoBz := new(bytes.Buffer)
			if err := json.Compact(contractInfoBz, contractInfo); err != nil {
				return err
			}

			msg := types.NewMsgMCOStore(contractInfoBz.Bytes(), cliCtx.FromAddress)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}
