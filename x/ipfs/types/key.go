package types

import (
	"github.com/ipfs/go-cid"
	dshelp "github.com/ipfs/go-ipfs-ds-help"
)

const (
	// ModuleName is the name of the module
	ModuleName = "ipfs"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querierer msgs
	QuerierRoute = ModuleName
)

// Keys for ipfs store
// Items are stored with the following key: values
//
// - 0x00<datastoreKey_Bytes>: Block
var (
	BlockKeyPrefix = []byte{0x00}
)

func GetBlockKey(c cid.Cid) []byte {
	return append(BlockKeyPrefix, dshelp.MultihashToDsKey(c.Hash()).Bytes()...)
}