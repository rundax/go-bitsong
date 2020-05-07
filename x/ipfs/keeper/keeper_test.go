package keeper

import (
	_ "fmt"
	blocks "github.com/ipfs/go-block-format"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKeeper_Put(t *testing.T) {
	ctx, k := SetupTestInput()

	data := []byte("hello world")
	block := k.Put(ctx, blocks.NewBlock(data))
	require.Equal(t, data, block.RawData())
}

func TestKeeper_Get(t *testing.T) {
	ctx, k := SetupTestInput()

	data := []byte("hello world")
	blockSaved := k.Put(ctx, blocks.NewBlock(data))

	block, found := k.Get(ctx, blockSaved.Cid())
	require.True(t, found)
	require.Equal(t, data, block.RawData())
	require.True(t, blockSaved.Cid().Equals(block.Cid()))
}
