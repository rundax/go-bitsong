package keeper

import (
	"fmt"
	"github.com/bitsongofficial/go-bitsong/x/mpeg21/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case types.MsgMpeg21StoreMCO:
			return handleMsgMpeg21StoreMCO(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized content message type: %T", msg.Type())
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgMpeg21StoreMCO(ctx sdk.Context, keeper Keeper, msg types.MsgMpeg21StoreMCO) (*sdk.Result, error) {
	contract, err := types.NewContract(msg.ContractInfo)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	contractID, err := keeper.Store(ctx, contract)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	/*for _, entity := range msg.Entities {
		// mint nft
		// TODO: convert with standard nft module
		coin := sdk.Coin{
			Denom:  track.ToCoinDenom(),
			Amount: entity.Shares, // TODO: entity shares must be > 0
		}
		if err := keeper.MintAndSend(ctx, coin, entity.Address); err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
	}*/

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMpeg21StoreMCO,
			sdk.NewAttribute(types.AttributeKeyContractID, fmt.Sprintf("%s", contractID)),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
