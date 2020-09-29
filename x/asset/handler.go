package asset

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/modules/incubator/nft"
)

func AssetHandler(keeper nft.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case nft.MsgMintNFT:
			return handleMsgMintNFT(ctx, msg, keeper)
		default:
			errMsg := fmt.Sprintf("unrecognized nft message type: %T", msg.Type())
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgMintNFT(ctx sdk.Context, msg nft.MsgMintNFT, k nft.Keeper) (*sdk.Result, error) {
	return nft.HandleMsgMintNFT(ctx, msg, k)
}
