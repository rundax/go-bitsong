package keeper

import (
	"encoding/hex"
	"github.com/bitsongofficial/go-bitsong/x/track/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"
	"time"
)

var (
	mockTitle               = "The Show Must Go On"
	mockIpfs                = "/ipfs/QmWWQSuPMS6aXCbZKpEjPHPUZN2NjB3YrhFTHsV4X3vb2t"
	mockRightHolder1        = types.NewRightHolder(sdk.AccAddress(crypto.AddressHash([]byte("rightHolder1"))), 100)
	mockRightsHoldersSingle = types.RightsHolders{
		mockRightHolder1,
	}
	mockRewards = types.TrackRewards{
		Users:     10,
		Playlists: 10,
	}
	mockOwner         = sdk.AccAddress(crypto.AddressHash([]byte("owner")))
	mockHexDecoded, _ = hex.DecodeString("B0FA2953B126722264F67828AF7443144C85D867")
	mockTrackAddr1    = crypto.Address(mockHexDecoded)
	mockTrack         = types.Track{
		Path:          mockIpfs,
		Rewards:       mockRewards,
		RightsHolders: mockRightsHoldersSingle,
		CreatedAt:     time.Time{},
		Owner:         nil,
	}
)

func SetupTestInput() (sdk.Context, Keeper) {
	// define store keys
	trackKey := sdk.NewKVStoreKey("track")

	// create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	ms.MountStoreWithDB(trackKey, sdk.StoreTypeIAVL, memDB)
	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	// create a Cdc and a context
	cdc := testCodec()
	ctx := sdk.NewContext(ms, abci.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())

	return ctx, NewKeeper(cdc, trackKey)
}

func testCodec() *codec.Codec {
	var cdc = codec.New()

	// register the different types
	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	types.RegisterCodec(cdc)

	cdc.Seal()
	return cdc
}
