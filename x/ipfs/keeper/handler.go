package keeper

import (
	"fmt"
	"github.com/bitsongofficial/go-bitsong/x/ipfs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	blocks "github.com/ipfs/go-block-format"
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case types.MsgPut:
			return handleMsgPut(ctx, keeper, msg)

		default:
			errMsg := fmt.Sprintf("unrecognized ipfs message type: %T", msg.Type())
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgPut(ctx sdk.Context, keeper Keeper, msg types.MsgPut) (*sdk.Result, error) {
	block := keeper.Put(ctx, blocks.NewBlock(msg.Data))

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeDataAdded,
			sdk.NewAttribute(types.AttributeKeyCid, block.Cid().String()),
		),
	)

	return &sdk.Result{
		Events: ctx.EventManager().Events().ToABCIEvents(),
	}, nil
}
