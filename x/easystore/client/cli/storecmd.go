package cli

import (
	//"errors"
	"strings"
//	"bufio"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	client "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/xuyp1991/cosaccount/x/easystore/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	//"github.com/cosmos/cosmos-sdk/client/utils"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

// GetBroadcastCommand returns the tx broadcast command.
func GeteasystoreCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "store [name] [value]",
		Short: "store a value on chain",
		Long: strings.TrimSpace(`just store a value for a simple test
`),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
		//	inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(client.GetTxEncoder(cdc))

			// if cliCtx.Offline {
			// 	return errors.New("cannot broadcast tx during offline mode")
			// }

			msgStoreData := types.NewMsgSetStore(args[0],args[1],cliCtx.GetFromAddress())
			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msgStoreData})
		},
	}

	return flags.PostCommands(cmd)[0]
}
