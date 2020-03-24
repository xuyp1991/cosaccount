package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "easystore"
	StoreKey = ModuleName
	RouterKey = ModuleName
	keyCommunityTax = "just for test"
)

// Storedata 
type Storedata struct {
	Value string         `protobuf:"bytes,1,opt,name=value,proto3" json:"value"`
	Owner sdk.AccAddress `protobuf:"bytes,1,opt,name=value,proto3" json:"owner"`
}

// Returns a new Whois with the minprice as the price
func NewStoredata() Storedata {
	return Storedata{
	}
}