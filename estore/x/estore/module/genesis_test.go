package estore_test

import (
	"testing"

	keepertest "estore/testutil/keeper"
	"estore/testutil/nullify"
	estore "estore/x/estore/module"
	"estore/x/estore/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EstoreKeeper(t)
	estore.InitGenesis(ctx, k, genesisState)
	got := estore.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
