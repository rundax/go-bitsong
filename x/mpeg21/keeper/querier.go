package keeper

import (
	"fmt"
	"github.com/bitsongofficial/go-bitsong/x/mpeg21/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryID:
			return queryContractByID(ctx, path[1:], req, keeper)
		case types.QueryContracts:
			return queryContracts(ctx, req, keeper)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown contract query endpoint")
		}
	}
}

func queryContracts(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	var params types.QueryContractsParams
	err := keeper.cdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	contracts := keeper.GetContractsPaginated(ctx, params)
	bz, err := codec.MarshalJSONIndent(keeper.cdc, contracts)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func queryContractByID(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	if path[0] == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("unknown contractID %s", path[0]))
	}

	// TODO: if path[0] is null or empty
	contract, found := keeper.GetContract(ctx, path[0])
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("contractID %s not found", path[0]))
	}

	bz, err := codec.MarshalJSONIndent(keeper.cdc, contract)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	return bz, nil
}
