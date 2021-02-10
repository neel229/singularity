package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateStock(t *testing.T) {
	arg := CreateStockParams{
		Ticker:  "David",
		Details: "Token of David Dobrik",
	}

	stock, err := testQueries.CreateStock(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, stock)

	require.Equal(t, arg.Ticker, stock.Ticker)
	require.Equal(t, arg.Details, stock.Details)

	require.NotZero(t, stock.ID)
}
