package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strings"
)

/****
 * Track Msg
 ***/

// Track messages types and routes
const (
	TypeMsgCreate = "create"
)

/****************************************
 * MsgCreate
 ****************************************/

var _ sdk.Msg = MsgCreate{}

// MsgCreateTrack defines Create message
type MsgCreate struct {
	Path          string         `json:"path" yaml:"path"`
	Rewards       TrackRewards   `json:"rewards" yaml:"rewards"`
	RightsHolders RightsHolders  `json:"rights_holders" yaml:"rights_holders"`
	Owner         sdk.AccAddress `json:"owner" yaml:"owner"`
}

func NewMsgCreate(path string, rewards TrackRewards, rightsHolders RightsHolders, owner sdk.AccAddress) MsgCreate {
	return MsgCreate{
		Path:          path,
		Rewards:       rewards,
		RightsHolders: rightsHolders,
		Owner:         owner,
	}
}

//nolint
func (msg MsgCreate) Route() string { return RouterKey }
func (msg MsgCreate) Type() string  { return TypeMsgCreate }

// ValidateBasic
func (msg MsgCreate) ValidateBasic() error {
	if len(strings.TrimSpace(msg.Path)) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("track path cannot be empty"))
	}

	if err := msg.Rewards.Validate(); err != nil {
		return err
	}

	if err := msg.RightsHolders.Validate(); err != nil {
		return err
	}

	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid owner: %s", msg.Owner.String()))
	}

	return nil
}

// String MsgCreate
func (msg MsgCreate) String() string {
	return fmt.Sprintf(`Create Message:
Path: %s
Rewards - %s
Rights Holders
%s
Owner: %s
`, msg.Path, msg.Rewards.String(), msg.RightsHolders.String(), msg.Owner.String())
}

// GetSignBytes encodes the message for signing
func (msg MsgCreate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
