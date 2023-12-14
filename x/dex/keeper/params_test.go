package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "xchange/testutil/keeper"
	"xchange/x/dex/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DexKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
