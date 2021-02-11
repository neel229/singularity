package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomCreatorStock(t *testing.T) CreatorStock {
	arg := CreateCreatorStockParams{
		CreatorID: 1,
		StockID:   1,
	}
	creatorStock, err := testQueries.CreateCreatorStock(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, creatorStock)

	require.Equal(t, arg.CreatorID, creatorStock.CreatorID)
	require.Equal(t, arg.StockID, creatorStock.StockID)

	require.NotZero(t, creatorStock.ID)

	return creatorStock
}

func TestCreateCreatorStock(t *testing.T) {
	createRandomCreatorStock(t)
}

func TestGetCreatorStock(t *testing.T) {
	creatorStock := createRandomCreatorStock(t)
	creatorStock1, err := testQueries.GetCreatorStock(context.Background(), creatorStock.ID)
	require.NoError(t, err)
	require.NotEmpty(t, creatorStock1)

	require.Equal(t, creatorStock.CreatorID, creatorStock1.CreatorID)
	require.Equal(t, creatorStock.StockID, creatorStock1.StockID)

	require.NotZero(t, creatorStock1.ID)
}

func TestListCreatorStocks(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCreatorStock(t)
	}

	arg := ListCreatorStocksParams{
		Limit:  5,
		Offset: 5,
	}

	creatorStocks, err := testQueries.ListCreatorStocks(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, creatorStocks, 5)

	for _, creatorStock := range creatorStocks {
		require.NotEmpty(t, creatorStock)
	}
}
