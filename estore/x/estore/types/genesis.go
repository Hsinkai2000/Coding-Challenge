package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ProductList: []Product{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in product
	productIdMap := make(map[uint64]bool)
	productCount := gs.GetProductCount()
	for _, elem := range gs.ProductList {
		if _, ok := productIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for product")
		}
		if elem.Id >= productCount {
			return fmt.Errorf("product id should be lower or equal than the last id")
		}
		productIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
