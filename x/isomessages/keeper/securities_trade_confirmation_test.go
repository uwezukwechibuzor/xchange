package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "xchange/testutil/keeper"
	"xchange/testutil/nullify"
	"xchange/x/isomessages/keeper"
	"xchange/x/isomessages/types"
)

func createNSecuritiesTradeConfirmation(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SecuritiesTradeConfirmation {
	items := make([]types.SecuritiesTradeConfirmation, n)
	for i := range items {
		items[i].Id = keeper.AppendSecuritiesTradeConfirmation(ctx, items[i])
	}
	return items
}

func TestSecuritiesTradeConfirmationGet(t *testing.T) {
	keeper, ctx := keepertest.IsomessagesKeeper(t)
	items := createNSecuritiesTradeConfirmation(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetSecuritiesTradeConfirmation(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestSecuritiesTradeConfirmationRemove(t *testing.T) {
	keeper, ctx := keepertest.IsomessagesKeeper(t)
	items := createNSecuritiesTradeConfirmation(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSecuritiesTradeConfirmation(ctx, item.Id)
		_, found := keeper.GetSecuritiesTradeConfirmation(ctx, item.Id)
		require.False(t, found)
	}
}

func TestSecuritiesTradeConfirmationGetAll(t *testing.T) {
	keeper, ctx := keepertest.IsomessagesKeeper(t)
	items := createNSecuritiesTradeConfirmation(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSecuritiesTradeConfirmation(ctx)),
	)
}

func TestSecuritiesTradeConfirmationCount(t *testing.T) {
	keeper, ctx := keepertest.IsomessagesKeeper(t)
	items := createNSecuritiesTradeConfirmation(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetSecuritiesTradeConfirmationCount(ctx))
}
