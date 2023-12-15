package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"xchange/x/isomessages/types"
)

func (k Keeper) SecuritiesTradeConfirmationAll(goCtx context.Context, req *types.QueryAllSecuritiesTradeConfirmationRequest) (*types.QueryAllSecuritiesTradeConfirmationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var securitiesTradeConfirmations []types.SecuritiesTradeConfirmation
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	securitiesTradeConfirmationStore := prefix.NewStore(store, types.KeyPrefix(types.SecuritiesTradeConfirmationKey))

	pageRes, err := query.Paginate(securitiesTradeConfirmationStore, req.Pagination, func(key []byte, value []byte) error {
		var securitiesTradeConfirmation types.SecuritiesTradeConfirmation
		if err := k.cdc.Unmarshal(value, &securitiesTradeConfirmation); err != nil {
			return err
		}

		securitiesTradeConfirmations = append(securitiesTradeConfirmations, securitiesTradeConfirmation)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSecuritiesTradeConfirmationResponse{SecuritiesTradeConfirmation: securitiesTradeConfirmations, Pagination: pageRes}, nil
}

func (k Keeper) SecuritiesTradeConfirmation(goCtx context.Context, req *types.QueryGetSecuritiesTradeConfirmationRequest) (*types.QueryGetSecuritiesTradeConfirmationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	securitiesTradeConfirmation, found := k.GetSecuritiesTradeConfirmation(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetSecuritiesTradeConfirmationResponse{SecuritiesTradeConfirmation: securitiesTradeConfirmation}, nil
}
