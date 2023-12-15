package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"xchange/x/isomessages/types"
)

// GetSecuritiesTradeConfirmationCount get the total number of securitiesTradeConfirmation
func (k Keeper) GetSecuritiesTradeConfirmationCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SecuritiesTradeConfirmationCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSecuritiesTradeConfirmationCount set the total number of securitiesTradeConfirmation
func (k Keeper) SetSecuritiesTradeConfirmationCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SecuritiesTradeConfirmationCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSecuritiesTradeConfirmation appends a securitiesTradeConfirmation in the store with a new id and update the count
func (k Keeper) AppendSecuritiesTradeConfirmation(
	ctx sdk.Context,
	securitiesTradeConfirmation types.SecuritiesTradeConfirmation,
) uint64 {
	// Create the securitiesTradeConfirmation
	count := k.GetSecuritiesTradeConfirmationCount(ctx)

	// Set the ID of the appended value
	securitiesTradeConfirmation.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecuritiesTradeConfirmationKey))
	appendedValue := k.cdc.MustMarshal(&securitiesTradeConfirmation)
	store.Set(GetSecuritiesTradeConfirmationIDBytes(securitiesTradeConfirmation.Id), appendedValue)

	// Update securitiesTradeConfirmation count
	k.SetSecuritiesTradeConfirmationCount(ctx, count+1)

	return count
}

// SetSecuritiesTradeConfirmation set a specific securitiesTradeConfirmation in the store
func (k Keeper) SetSecuritiesTradeConfirmation(ctx sdk.Context, securitiesTradeConfirmation types.SecuritiesTradeConfirmation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecuritiesTradeConfirmationKey))
	b := k.cdc.MustMarshal(&securitiesTradeConfirmation)
	store.Set(GetSecuritiesTradeConfirmationIDBytes(securitiesTradeConfirmation.Id), b)
}

// GetSecuritiesTradeConfirmation returns a securitiesTradeConfirmation from its id
func (k Keeper) GetSecuritiesTradeConfirmation(ctx sdk.Context, id uint64) (val types.SecuritiesTradeConfirmation, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecuritiesTradeConfirmationKey))
	b := store.Get(GetSecuritiesTradeConfirmationIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSecuritiesTradeConfirmation removes a securitiesTradeConfirmation from the store
func (k Keeper) RemoveSecuritiesTradeConfirmation(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecuritiesTradeConfirmationKey))
	store.Delete(GetSecuritiesTradeConfirmationIDBytes(id))
}

// GetAllSecuritiesTradeConfirmation returns all securitiesTradeConfirmation
func (k Keeper) GetAllSecuritiesTradeConfirmation(ctx sdk.Context) (list []types.SecuritiesTradeConfirmation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SecuritiesTradeConfirmationKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SecuritiesTradeConfirmation
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSecuritiesTradeConfirmationIDBytes returns the byte representation of the ID
func GetSecuritiesTradeConfirmationIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetSecuritiesTradeConfirmationIDFromBytes returns ID in uint64 format from a byte array
func GetSecuritiesTradeConfirmationIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
