package cli

import (
	"bufio"
	"fmt"
	"github.com/bitsongofficial/go-bitsong/x/ipfs/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	chunker "github.com/ipfs/go-ipfs-chunker"
	"github.com/spf13/cobra"
	"io"
	"os"
)

func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	trackTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	trackTxCmd.AddCommand(flags.PostCommands(
		GetCmdPut(cdc),
	)...)

	return trackTxCmd
}

func GetCmdPut(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "put [file]",
		Short: "Put a new file",
		Long:  fmt.Sprintf(`%s`, version.ClientName),
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			/*data, err := ioutil.ReadFile(args[0])
			if err != nil {
				return err
			}*/

			file, _ := os.Open(args[0])
			chnk, err := chunker.FromString(bufio.NewReader(file), "size-262144")
			if err != nil {
				return err
			}

			for {
				chunk, err := chnk.NextBytes()
				if err != nil {
					if err == io.EOF {
						break
					}
					panic(err)
				}
				//res = res + uint64(len(chunk))
				err = authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{types.NewMsgPut(chunk, cliCtx.FromAddress)})
				if err != nil {
					panic(err)
				}
			}


			//msg := types.NewMsgPut(data, cliCtx.FromAddress)
			//return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, msgs)
			return nil
		},
	}

	return cmd
}