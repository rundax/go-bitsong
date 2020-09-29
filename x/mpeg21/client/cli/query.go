package cli

import (
	"fmt"
	"github.com/bitsongofficial/go-bitsong/x/mpeg21/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

const (
	flagPage  = "page"
	flagLimit = "limit"
)

// GetQueryCmd returns the cli query commands
func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	// Group content queries under a subcommand
	contentQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	contentQueryCmd.AddCommand(
		flags.GetCommands(
			GetCmdAll(cdc),
			GetCmdID(cdc),
		)...,
	)

	return contentQueryCmd
}

func GetCmdAll(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mco-all",
		Short: "Get all MCO Contracts",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Get all MCO Contracts.
Example:
$ %s query mpeg21 mco all
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			page := viper.GetInt(flagPage)
			limit := viper.GetInt(flagLimit)
			params := types.DefaultQueryContractsParams(page, limit)

			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryContracts)
			res, _, err := cliCtx.QueryWithData(route, bz)
			if err != nil {
				return err
			}

			var contracts []types.Contract
			if err := cdc.UnmarshalJSON(res, &contracts); err != nil {
				return err
			}

			return cliCtx.PrintOutput(contracts)
		},
	}

	cmd.Flags().Int(flagLimit, 100, "pagination limit of tracks to query for")
	cmd.Flags().Int(flagPage, 1, "pagination page of tracks to to query for")

	return cmd
}

func GetCmdID(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "mco-id [id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query a MCO by contractID",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query a MCO by contractID.
Example:
$ %s query mco [id]
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			id := args[0]

			route := fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryID, id)
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				fmt.Printf("Could not find contractID %s \n", id)
				return nil
			}

			var contract types.Contract
			cdc.MustUnmarshalJSON(res, &contract)
			return cliCtx.PrintOutput(contract)
		},
	}
}
