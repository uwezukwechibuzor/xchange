package keeper

import (
	"xchange/x/dex/types"
)

var _ types.QueryServer = Keeper{}
