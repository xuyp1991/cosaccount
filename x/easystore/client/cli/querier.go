package cli

import (
	"github.com/spf13/cobra"
	"github.com/cosmos/cosmos-sdk/codec"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/xuyp1991/cosaccount/x/easystore/types"
)

// GetCmdResolveName queries information about a name
func GetCmdValue(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "values [name]",
		Short: "query name store value",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			name := args[0]

			res, _,err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/values/%s", queryRoute, name), nil)
			if err != nil {
				fmt.Printf("could not resolve name - %s \n", string(name))
				return nil
			}

			var out types.QueryResResolve
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}