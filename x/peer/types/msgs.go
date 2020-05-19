package types

import sdk "github.com/cosmos/cosmos-sdk/types"
import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

const (
	TypeMsgPut = "put"
)

/****************************************
 * MsgPut
 ****************************************/

var _ sdk.Msg = MsgPut{}

type MsgPut struct {
	From sdk.AccAddress `json:"from" yaml:"from"`
	Data []byte `json:"data" yaml:"data"`
}

func NewMsgPut(data []byte, from sdk.AccAddress) MsgPut {
	return MsgPut{
		Data: data,
		From: from,
	}
}

func (msg MsgPut) Route() string { return RouterKey }
func (msg MsgPut) Type() string  { return TypeMsgPut }

func (msg MsgPut) ValidateBasic() error {
	if len(msg.Data) > 200000000 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "file is to bigger")
	}

	return nil
}

func (msg MsgPut) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgPut) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}