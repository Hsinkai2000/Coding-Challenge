package estore

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"estore/x/estore/keeper"
	"estore/x/estore/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the product
	for _, elem := range genState.ProductList {
		k.SetProduct(ctx, elem)
	}

	// Set product count
	k.SetProductCount(ctx, genState.ProductCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ProductList = k.GetAllProduct(ctx)
	genesis.ProductCount = k.GetProductCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
