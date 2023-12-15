package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		SecuritiesTradeConfirmationList: []SecuritiesTradeConfirmation{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in securitiesTradeConfirmation
	securitiesTradeConfirmationIdMap := make(map[uint64]bool)
	securitiesTradeConfirmationCount := gs.GetSecuritiesTradeConfirmationCount()
	for _, elem := range gs.SecuritiesTradeConfirmationList {
		if _, ok := securitiesTradeConfirmationIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for securitiesTradeConfirmation")
		}
		if elem.Id >= securitiesTradeConfirmationCount {
			return fmt.Errorf("securitiesTradeConfirmation id should be lower or equal than the last id")
		}
		securitiesTradeConfirmationIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
