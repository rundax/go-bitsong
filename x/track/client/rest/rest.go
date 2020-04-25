package rest

import (
	"github.com/bitsongofficial/go-bitsong/x/track/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers track-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	registerQueryRoutes(cliCtx, r)
	registerTxRoutes(cliCtx, r)
}

type CreateTrackReq struct {
	BaseReq       rest.BaseReq        `json:"base_req"`
	Path          string              `json:"path"`
	Rewards       types.TrackRewards  `json:"rewards"`
	RightsHolders types.RightsHolders `json:"rights_holders"`
	Owner         sdk.AccAddress      `json:"owner"`
}
