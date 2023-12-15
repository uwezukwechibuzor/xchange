package isomessages_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "xchange/testutil/keeper"
	"xchange/testutil/nullify"
	"xchange/x/isomessages"
	"xchange/x/isomessages/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SecuritiesTradeConfirmationList: []types.SecuritiesTradeConfirmation{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		SecuritiesTradeConfirmationCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IsomessagesKeeper(t)
	isomessages.InitGenesis(ctx, *k, genesisState)
	got := isomessages.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.SecuritiesTradeConfirmationList, got.SecuritiesTradeConfirmationList)
	require.Equal(t, genesisState.SecuritiesTradeConfirmationCount, got.SecuritiesTradeConfirmationCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
