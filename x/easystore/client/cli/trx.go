package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	easystore "github.com/xuyp1991/cosaccount/x/easystore/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        easystore.ModuleName,
		Short:                      "Auth transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(
		GeteasystoreCmd(cdc),
		// GetSignCommand(cdc),
	)
	return txCmd
}

// GetTxCmd returns the transaction commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        easystore.ModuleName,
		Short:                      "Auth transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(
		GetCmdValue(queryRoute,cdc),
		// GetSignCommand(cdc),
	)
	return txCmd
}