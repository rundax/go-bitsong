package keeper

import (
	"fmt"
	"github.com/bitsongofficial/go-bitsong/x/peer/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper of the track store
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      *codec.Codec
}

// NewKeeper creates a track keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
	keeper := Keeper{
		storeKey: key,
		cdc:      cdc,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) setBlock(ctx sdk.Context, block *blocks.BasicBlock) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetBlockKey(block.Cid()), block.RawData())
	return
}

func (k Keeper) getBlock(ctx sdk.Context, c cid.Cid) (block *blocks.BasicBlock, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetBlockKey(c))
	if bz == nil {
		return
	}
	return blocks.NewBlock(bz), true
}

func (k Keeper) Put(ctx sdk.Context, block *blocks.BasicBlock) *blocks.BasicBlock {
	k.setBlock(ctx, block)
	return block
}

func (k Keeper) Get(ctx sdk.Context, c cid.Cid) (*blocks.BasicBlock, bool) {
	return k.getBlock(ctx, c)
}
