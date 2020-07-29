// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/deposit/types
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/deposit/keeper
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/deposit/querier
package deposit

import (
	"github.com/sentinel-official/hub/x/dvpn/deposit/keeper"
	"github.com/sentinel-official/hub/x/dvpn/deposit/querier"
	"github.com/sentinel-official/hub/x/dvpn/deposit/types"
)

const (
	Codespace     = types.Codespace
	ModuleName    = types.ModuleName
	QuerierRoute  = types.QuerierRoute
	QueryDeposit  = types.QueryDeposit
	QueryDeposits = types.QueryDeposits
)

var (
	// functions aliases
	RegisterCodec                 = types.RegisterCodec
	ErrorMarshal                  = types.ErrorMarshal
	ErrorUnmarshal                = types.ErrorUnmarshal
	ErrorUnknownQueryType         = types.ErrorUnknownQueryType
	ErrorInsufficientDepositFunds = types.ErrorInsufficientDepositFunds
	ErrorDepositDoesNotExist      = types.ErrorDepositDoesNotExist
	NewGenesisState               = types.NewGenesisState
	DefaultGenesisState           = types.DefaultGenesisState
	DepositKey                    = types.DepositKey
	NewQueryDepositParams         = types.NewQueryDepositParams
	NewKeeper                     = keeper.NewKeeper
	Querier                       = querier.Querier

	// variable aliases
	ModuleCdc        = types.ModuleCdc
	RouterKey        = types.RouterKey
	StoreKey         = types.StoreKey
	DepositKeyPrefix = types.DepositKeyPrefix
)

type (
	Deposit            = types.Deposit
	Deposits           = types.Deposits
	GenesisState       = types.GenesisState
	QueryDepositParams = types.QueryDepositParams
	Keeper             = keeper.Keeper
)