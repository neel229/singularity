package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomStock(t *testing.T) Stock {
	arg := CreateStockParams{
		Ticker:  "LPT",
		Details: "Token of Logan Paul",
	}

	stock, err := testQueries.CreateStock(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, stock)

	require.Equal(t, arg.Ticker, stock.Ticker)
	require.Equal(t, arg.Details, stock.Details)
	require.NotZero(t, stock.ID)

	return stock
}

func TestCreateStock(t *testing.T) {
	createRandomStock(t)
}

func TestGetStock(t *testing.T) {
	stock := createRandomStock(t)
	stock1, err := testQueries.GetStock(context.Background(), stock.ID)
	require.NoError(t, err)
	require.NotEmpty(t, stock1)

	require.Equal(t, stock.Ticker, stock.Ticker)
	require.Equal(t, stock.Details, stock.Details)
	require.NotZero(t, stock1.ID)
}

func TestListStocks(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomStock(t)
	}

	arg := ListStocksParams{
		Limit:  5,
		Offset: 5,
	}

	stocks, err := testQueries.ListStocks(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, stocks, 5)

	for _, stock := range stocks {
		require.NotEmpty(t, stock)
	}
}

func TestUpdateStock(t *testing.T) {
	stock := createRandomStock(t)
	arg := UpdateStockParams{
		ID:      stock.ID,
		Details: "xvnxvoiwer",
	}
	err := testQueries.UpdateStock(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, stock.Details)
}

func TestDeleteStock(t *testing.T) {
	stock := createRandomStock(t)
	err := testQueries.DeleteStock(context.Background(), stock.ID)
	require.NoError(t, err)
	stock1, err := testQueries.GetStock(context.Background(), stock.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, stock1)
}
