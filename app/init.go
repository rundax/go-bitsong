package app

import (
	btsg "github.com/bitsongofficial/go-bitsong/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/staking"
)

func Init() {
	staking.DefaultGenesisState = stakingGenesisState
	mint.DefaultGenesisState = mintGenesisState
	gov.DefaultGenesisState = govGenesisState
	crisis.DefaultGenesisState = crisisGenesisState
}

func stakingGenesisState() staking.GenesisState {
	return staking.GenesisState{
		Params: staking.NewParams(
			staking.DefaultUnbondingTime,
			staking.DefaultMaxValidators,
			staking.DefaultMaxEntries,
			0,
			btsg.BondDenom,
		),
	}
}

func mintGenesisState() mint.GenesisState {
	params := mint.DefaultParams()

	return mint.GenesisState{
		Params: mint.NewParams(
			btsg.BondDenom,
			params.InflationRateChange,
			params.InflationMax,
			params.InflationMin,
			params.GoalBonded,
			params.BlocksPerYear,
		),
	}
}

func govGenesisState() gov.GenesisState {
	return gov.NewGenesisState(
		govtypes.DefaultStartingProposalID,
		gov.NewDepositParams(
			sdk.NewCoins(sdk.NewCoin(btsg.BondDenom, govtypes.DefaultMinDepositTokens)),
			govtypes.DefaultPeriod,
		),
		govtypes.DefaultVotingParams(),
		govtypes.DefaultTallyParams(),
	)
}

func crisisGenesisState() crisis.GenesisState {
	return crisis.NewGenesisState(sdk.NewCoin(btsg.BondDenom, sdk.NewInt(1000)))
}