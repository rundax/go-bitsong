package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"
	"testing"
	"time"
)

var (
	mockEmptyString         = ""
	mockIpfs                = "/ipfs/QmWWQSuPMS6aXCbZKpEjPHPUZN2NjB3YrhFTHsV4X3vb2t"
	mockRightHolder1        = NewRightHolder(sdk.AccAddress(crypto.AddressHash([]byte("rightHolder1"))), 100)
	mockRightsHoldersSingle = RightsHolders{
		mockRightHolder1,
	}
	mockRewards = TrackRewards{
		Users:     10,
		Playlists: 10,
	}
	mockOwner = sdk.AccAddress(crypto.AddressHash([]byte("owner")))
	mockTrack = Track{
		Path:          mockIpfs,
		Rewards:       mockRewards,
		RightsHolders: mockRightsHoldersSingle,
		CreatedAt:     time.Time{},
		Owner:         mockOwner,
	}
	mockTrackOwnerNil = Track{
		Path:          mockIpfs,
		Rewards:       mockRewards,
		RightsHolders: mockRightsHoldersSingle,
		CreatedAt:     time.Time{},
		Owner:         nil,
	}
)

var mockMsgCreate = NewMsgCreate(
	mockIpfs,
	mockRewards,
	mockRightsHoldersSingle,
	mockOwner,
)

func TestMsgCreate_Route(t *testing.T) {
	expected := "track"
	actual := mockMsgCreate.Route()
	require.Equal(t, expected, actual)
}

func TestMsgCreate_Type(t *testing.T) {
	expected := "create"
	actual := mockMsgCreate.Type()
	require.Equal(t, expected, actual)
}

func TestMsgCreate_ValidateBasic(t *testing.T) {
	_ = sdk.AccAddress(crypto.AddressHash([]byte("test")))

	// TODO: continue with more test
	tests := []struct {
		name  string
		msg   MsgCreate
		error error
	}{
		{
			name:  "Empty owner return error",
			msg:   NewMsgCreate(mockIpfs, mockRewards, mockRightsHoldersSingle, nil),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid owner: "),
		},
		{
			name:  "Empty path returns error if path is empty",
			msg:   NewMsgCreate(mockEmptyString, mockRewards, mockRightsHoldersSingle, mockOwner),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "track path cannot be empty"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			returnedError := test.msg.ValidateBasic()
			if test.error == nil {
				require.Nil(t, returnedError)
			} else {
				require.NotNil(t, returnedError)
				require.Equal(t, test.error.Error(), returnedError.Error())
			}
		})
	}

	err := mockMsgCreate.ValidateBasic()
	require.Nil(t, err)
}

func TestMsgCreate_GetSignBytes(t *testing.T) {
	tests := []struct {
		name        string
		msg         MsgCreate
		expSignJSON string
	}{
		{
			name:        "Message with no path",
			msg:         NewMsgCreate(mockEmptyString, mockRewards, mockRightsHoldersSingle, mockOwner),
			expSignJSON: `{"type":"go-bitsong/MsgCreateTrack","value":{"owner":"cosmos1fsgzj6t7udv8zhf6zj32mkqhcjcpv52ygswxa5","path":"","rewards":{"playlists":"10","users":"10"},"rights_holders":[{"address":"cosmos17zfanegzaj8shhzsrfncz6cz5ykvzr06yyww88","quota":"100"}]}}`,
		},
		{
			name:        "Message with owner nil",
			msg:         NewMsgCreate(mockEmptyString, mockRewards, mockRightsHoldersSingle, nil),
			expSignJSON: `{"type":"go-bitsong/MsgCreateTrack","value":{"owner":"","path":"","rewards":{"playlists":"10","users":"10"},"rights_holders":[{"address":"cosmos17zfanegzaj8shhzsrfncz6cz5ykvzr06yyww88","quota":"100"}]}}`,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expSignJSON, string(test.msg.GetSignBytes()))
		})
	}
}

func TestMsgCreate_GetSigners(t *testing.T) {
	expected := mockMsgCreate.Owner
	actual := mockMsgCreate.GetSigners()
	require.Equal(t, expected, actual[0])
	require.Equal(t, 1, len(actual))
}
