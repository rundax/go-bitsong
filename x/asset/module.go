package asset

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/modules/incubator/nft"
)

type AssetModule struct {
	nft.AppModule
	k nft.Keeper
}

func NewAssetModule(appModule nft.AppModule, keeper nft.Keeper) AssetModule {
	return AssetModule{
		AppModule: appModule,
		k:         keeper,
	}
}

func (am AssetModule) NewHandler() sdk.Handler {
	return AssetHandler(am.k)
}
