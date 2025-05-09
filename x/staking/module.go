package staking

import (
	"context"

	"cosmossdk.io/core/appmodule"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/staking/exported"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/xpladev/xpla/x/staking/keeper"
)

var (
	_ module.AppModule          = staking.AppModule{}
	_ module.HasABCIEndBlock    = AppModule{}
	_ appmodule.HasBeginBlocker = AppModule{}
)

type AppModule struct {
	staking.AppModule

	keeper *keeper.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec, keeper *keeper.Keeper, ak stakingtypes.AccountKeeper, bk stakingtypes.BankKeeper, ls exported.Subspace) AppModule {
	return AppModule{
		AppModule: staking.NewAppModule(cdc, keeper.Keeper, ak, bk, ls),
		keeper:    keeper,
	}
}

// BeginBlock returns the begin blocker for the staking module.
func (am AppModule) BeginBlock(ctx context.Context) error {
	return BeginBlocker(ctx, am.keeper)
}

// EndBlock returns the end blocker for the staking module. It returns no validator
// updates.
func (am AppModule) EndBlock(ctx context.Context) ([]abci.ValidatorUpdate, error) {
	return EndBlocker(ctx, am.keeper)
}
