package track

import (
	"encoding/hex"
	btsg "github.com/bitsongofficial/go-bitsong/types"
	"github.com/bitsongofficial/go-bitsong/x/track/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
	"testing"
	"time"
)

var (
	mockRightHolder1        = types.NewRightHolder(sdk.AccAddress(crypto.AddressHash([]byte("rightHolder1"))), 100)
	mockRightsHoldersSingle = types.RightsHolders{
		mockRightHolder1,
	}
	mockRewards = types.TrackRewards{
		Users:     10,
		Playlists: 10,
	}
	timeZone, _   = time.LoadLocation("UTC")
	mockDate      = time.Date(2020, 1, 1, 12, 00, 00, 000, timeZone)
	mockOwner     = sdk.AccAddress(crypto.AddressHash([]byte("owner")))
	tAddr, _      = hex.DecodeString("B0FA2953B126722264F67828AF7443144C85D867")
	mockTrackAddr = crypto.Address(tAddr)
	mockTracks    = types.Tracks{
		types.Track{
			Path:          "/ipfs/QmWWQSuPMS6aXCbZKpEjPHPUZN2NjB3YrhFTHsV4X3vb2t",
			Address:       mockTrackAddr,
			Rewards:       mockRewards,
			RightsHolders: mockRightsHoldersSingle,
			Totals: types.TrackTotals{
				Streams:  0,
				Rewards:  sdk.NewCoin(btsg.BondDenom, sdk.ZeroInt()),
				Accounts: 0,
			},
			CreatedAt: mockDate,
			Owner:     mockOwner,
		},
	}
)

type TestInput struct {
	Ctx         sdk.Context
	TrackKeeper Keeper
}

func newTestCodec() *codec.Codec {
	cdc := codec.New()

	types.RegisterCodec(cdc)

	return cdc
}

func CreateTestInput(t *testing.T) TestInput {
	keyTrack := sdk.NewKVStoreKey(types.StoreKey)

	cdc := newTestCodec()
	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ctx := sdk.NewContext(ms, abci.Header{Time: time.Now().UTC()}, false, log.NewNopLogger())

	ms.MountStoreWithDB(keyTrack, sdk.StoreTypeIAVL, db)
	require.NoError(t, ms.LoadLatestVersion())

	trackKeeper := NewKeeper(cdc, keyTrack)

	return TestInput{ctx, trackKeeper}
}

// TODO: improve tests
func TestInitGenesis(t *testing.T) {
	tests := []struct {
		name              string
		importTracks      types.Tracks
		importLastTrackID uint64
		expLastTrackID    uint64
		expTracks         types.Tracks
		expError          error
	}{
		{
			name:           "Expected error if initial track ID",
			importTracks:   mockTracks,
			expLastTrackID: uint64(1),
			expTracks:      mockTracks,
		},
		{
			name:              "Expected no error with lastrackid 1 and mocktracks",
			importLastTrackID: 1,
			importTracks:      mockTracks,
			expLastTrackID:    uint64(1),
			expTracks:         mockTracks,
		},
		{
			name:              "Expected error if expLasTrackID is different from storedTracks",
			importLastTrackID: 1,
			importTracks:      mockTracks,
			expLastTrackID:    uint64(2),
			expTracks:         mockTracks,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			input := CreateTestInput(t)

			var lastTrackID uint64

			if test.importLastTrackID == 0 {
				_, err := input.TrackKeeper.GetLastTrackID(input.Ctx)
				require.Error(t, err, "initial track ID hasn't been set")

				lastTrackID = 1
			} else {
				lastTrackID = test.importLastTrackID
			}

			genesisState := types.NewGenesisState(lastTrackID, test.importTracks)

			InitGenesis(input.Ctx, input.TrackKeeper, genesisState)

			actualLastTrackID, err := input.TrackKeeper.GetLastTrackID(input.Ctx)
			require.NoError(t, err)

			if test.importLastTrackID == test.expLastTrackID {
				require.Equal(t, test.expLastTrackID, actualLastTrackID)
			}

			actualTracks := input.TrackKeeper.GetTracks(input.Ctx)
			require.Equal(t, test.expTracks, actualTracks)
		})
	}
}

// TODO: Add more tests
func TestExportGenesis(t *testing.T) {
	tests := []struct {
		name         string
		genesisState GenesisState
	}{
		{
			name:         "Expected equal genesis",
			genesisState: NewGenesisState(1, mockTracks),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			input := CreateTestInput(t)

			InitGenesis(input.Ctx, input.TrackKeeper, test.genesisState)
			exported := ExportGenesis(input.Ctx, input.TrackKeeper)

			require.Equal(t, test.genesisState, exported)
		})
	}
}
