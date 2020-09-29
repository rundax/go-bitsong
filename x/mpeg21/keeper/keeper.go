package keeper

import (
	"fmt"
	"github.com/bitsongofficial/go-bitsong/x/mpeg21/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/libs/log"
	"sort"
)

type Keeper struct {
	storeKey sdk.StoreKey
	cdc      *codec.Codec
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
	keeper := Keeper{
		storeKey: key,
		cdc:      cdc,
	}

	return keeper
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetContract(ctx sdk.Context, contractID string) (contract types.Contract, ok bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetContractKey(contractID))
	if bz == nil {
		return
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &contract)
	return contract, true
}

func (k Keeper) SetContract(ctx sdk.Context, contract *types.Contract) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(&contract)
	store.Set(types.GetContractKey(contract.ID), bz)
}

func (k Keeper) IterateContracts(ctx sdk.Context, fn func(contract types.Contract) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ContractKeyPrefix)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var contract types.Contract
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iterator.Value(), &contract)
		if fn(contract) {
			break
		}
	}
}

func (k Keeper) GetContracts(ctx sdk.Context) []types.Contract {
	var contracts []types.Contract
	k.IterateContracts(ctx, func(contract types.Contract) (stop bool) {
		contracts = append(contracts, contract)
		return false
	})
	return contracts
}

func (k Keeper) GetContractsPaginated(ctx sdk.Context, params types.QueryContractsParams) []types.Contract {
	var contracts []types.Contract
	k.IterateContracts(ctx, func(contract types.Contract) (stop bool) {
		contracts = append(contracts, contract)
		return false
	})

	sort.Slice(contracts, func(i, j int) bool {
		a, b := contracts[i], contracts[j]
		return a.ID > b.ID
	})

	page := params.Page
	if page == 0 {
		page = 1
	}
	start, end := client.Paginate(len(contracts), page, params.Limit, 100)
	if start < 0 || end < 0 {
		contracts = []types.Contract{}
	} else {
		contracts = contracts[start:end]
	}

	return contracts
}

func (k Keeper) Store(ctx sdk.Context, contract *types.Contract) (string, error) {
	_, ok := k.GetContract(ctx, contract.ID)
	if ok {
		return contract.ID, sdkerrors.Wrapf(types.ErrDuplicatedContract, "%s", contract.ID)
	}

	k.SetContract(ctx, contract)

	return contract.ID, nil
}
