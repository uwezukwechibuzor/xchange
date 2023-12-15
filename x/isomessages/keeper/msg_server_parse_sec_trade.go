package keeper

import (
	"context"

	"xchange/x/isomessages/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ParseSecTrade is a method of the msgServer type, implementing the ParseSecTrade
// RPC endpoint. It handles the parsing of a securities trade message and
// stores the relevant information in the application's state.
func (k msgServer) ParseSecTrade(goCtx context.Context, msg *types.MsgParseSecTrade) (*types.MsgParseSecTradeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	securitiesTradeConfirmation := types.SecuritiesTradeConfirmation{
		//Id: msg.TradConfId,
		TradConfId:   msg.TradConfId,
		TradConfDtTm: msg.TradConfDtTm,
		RltdOrdrId:   msg.RltdOrdrId,
		// Include other fields from msg as needed.
	}

	// set to store
	k.SetSecuritiesTradeConfirmation(ctx, securitiesTradeConfirmation)

	return &types.MsgParseSecTradeResponse{}, nil
}
