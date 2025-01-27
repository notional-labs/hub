package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	ModuleName   = "deposit"
	QuerierRoute = ModuleName
)

var (
	RouterKey = ModuleName
	StoreKey  = ModuleName
)

var (
	DepositKeyPrefix = []byte{0x10}
)

func DepositKey(addr sdk.AccAddress) []byte {
	return append(DepositKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}
