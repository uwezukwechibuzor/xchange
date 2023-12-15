package isomessages

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"xchange/x/isomessages/keeper"
	"xchange/x/isomessages/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the securitiesTradeConfirmation
	for _, elem := range genState.SecuritiesTradeConfirmationList {
		k.SetSecuritiesTradeConfirmation(ctx, elem)
	}

	// Set securitiesTradeConfirmation count
	k.SetSecuritiesTradeConfirmationCount(ctx, genState.SecuritiesTradeConfirmationCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.SecuritiesTradeConfirmationList = k.GetAllSecuritiesTradeConfirmation(ctx)
	genesis.SecuritiesTradeConfirmationCount = k.GetSecuritiesTradeConfirmationCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
