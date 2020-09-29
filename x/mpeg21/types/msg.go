package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgMpeg21StoreMCO = "mpeg21_store_mco"
)

var _ sdk.Msg = MsgMpeg21StoreMCO{}

type MsgMpeg21StoreMCO struct {
	ContractInfo []byte         `json:"contract_info" yaml:"contract_info"`
	Creator      sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgMCOStore(contractInfo []byte, creator sdk.AccAddress) MsgMpeg21StoreMCO {
	return MsgMpeg21StoreMCO{
		ContractInfo: contractInfo,
		Creator:      creator,
	}
}

func (msg MsgMpeg21StoreMCO) Route() string { return RouterKey }
func (msg MsgMpeg21StoreMCO) Type() string  { return TypeMsgMpeg21StoreMCO }

func (msg MsgMpeg21StoreMCO) ValidateBasic() error {
	if len(msg.ContractInfo) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "contract info cannot be empty")
	}

	/*if len(msg.ContractInfo) > MaxContractInfoLength {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "contract info too large")
	}*/

	return nil
}

func (msg MsgMpeg21StoreMCO) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgMpeg21StoreMCO) GetSigners() []sdk.AccAddress {
	// TODO: all party must sign
	return []sdk.AccAddress{msg.Creator}
}

func (msg MsgMpeg21StoreMCO) String() string {
	return fmt.Sprintf(`Msg Contract Store
TrackInfo: %s,
Creator: %s`,
		msg.ContractInfo, msg.Creator,
	)
}
