package functt

import (
//	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	//sdk "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/xuyp1991/cosaccount/x/easystore/types"

)

// // NewHandler returns a handler for "nameservice" type messages.
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types.MsgSetStore:
			return handleMsgSetStore(ctx, k, msg)
		default:
		//	errMsg := fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.Result{}
		}
	}
}

// Handle a message to buy name
func handleMsgSetStore(ctx sdk.Context, keeper Keeper, msg types.MsgSetStore) sdk.Result {
	keeper.Setvalue(ctx,msg.Name,msg.Value,msg.Owner)
	return sdk.Result{}
}