package keeper_test

import (
	"context"
	"testing"

	keepertest "estore/testutil/keeper"
	"estore/testutil/nullify"
	"estore/x/estore/keeper"
	"estore/x/estore/types"

	"github.com/stretchr/testify/require"
)

func createNProduct(keeper keeper.Keeper, ctx context.Context, n int) []types.Product {
	items := make([]types.Product, n)
	for i := range items {
		items[i].Id = keeper.AppendProduct(ctx, items[i])
	}
	return items
}

func TestProductGet(t *testing.T) {
	keeper, ctx := keepertest.EstoreKeeper(t)
	items := createNProduct(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetProduct(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestProductRemove(t *testing.T) {
	keeper, ctx := keepertest.EstoreKeeper(t)
	items := createNProduct(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProduct(ctx, item.Id)
		_, found := keeper.GetProduct(ctx, item.Id)
		require.False(t, found)
	}
}

func TestProductGetAll(t *testing.T) {
	keeper, ctx := keepertest.EstoreKeeper(t)
	items := createNProduct(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllProduct(ctx)),
	)
}

func TestProductCount(t *testing.T) {
	keeper, ctx := keepertest.EstoreKeeper(t)
	items := createNProduct(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetProductCount(ctx))
}
