package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateCreatorStock(t *testing.T) {
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
}
