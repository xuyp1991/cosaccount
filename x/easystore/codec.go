package functt

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/xuyp1991/cosaccount/x/easystore/types"
)

// RegisterCodec registers concrete types on wire codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(types.MsgSetStore{}, "cosaccount/SetStore", nil)
}

// module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}
