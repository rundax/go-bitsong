package types

type GenesisState struct {
	Contracts []Contract `json:"contracts"`
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState(contracts []Contract) GenesisState {
	return GenesisState{
		Contracts: contracts,
	}
}

// DefaultGenesisState - default GenesisState
func DefaultGenesisState() GenesisState {
	return GenesisState{
		Contracts: []Contract{},
	}
}

// ValidateGenesis validates the tracks genesis parameters
func ValidateGenesis(data GenesisState) error {
	for _, item := range data.Contracts {
		if err := item.Validate(); err != nil {
			return err
		}
	}

	return nil
}
