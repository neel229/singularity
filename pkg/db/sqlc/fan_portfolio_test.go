package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomFanPortfolio(t *testing.T) FanPortfolio {
	arg := CreateFanPortfolioParams{
		FanID:    1,
		StockID:  1,
		Quantity: "10000000.000000",
	}
	portfolio, err := testQueries.CreateFanPortfolio(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, portfolio)

	require.Equal(t, arg.FanID, portfolio.FanID)
	require.Equal(t, arg.StockID, portfolio.StockID)
	require.Equal(t, arg.Quantity, portfolio.Quantity)

	require.NotZero(t, portfolio.ID)

	return portfolio
}

func TestCreateFanPortfolio(t *testing.T) {
	createRandomFanPortfolio(t)
}

func TestGetFanPortfolio(t *testing.T) {
	portfolio := createRandomFanPortfolio(t)
	portfolio1, err := testQueries.GetFanPortfolio(context.Background(), portfolio.ID)
	require.NoError(t, err)
	require.NotEmpty(t, portfolio1)

	require.Equal(t, portfolio.FanID, portfolio1.FanID)
	require.Equal(t, portfolio.StockID, portfolio1.StockID)
	require.Equal(t, portfolio.Quantity, portfolio1.Quantity)

	require.NotZero(t, portfolio1.ID)
}

func TestGetPortfolioByFanID(t *testing.T) {
	portfolio := createRandomFanPortfolio(t)
	portfolio1, err := testQueries.GetPortfolioByFanID(context.Background(), portfolio.FanID)
	require.NoError(t, err)
	require.NotEmpty(t, portfolio1)

	require.Equal(t, portfolio.FanID, portfolio1.FanID)
	require.Equal(t, portfolio.StockID, portfolio1.StockID)
	require.Equal(t, portfolio.Quantity, portfolio1.Quantity)

	require.NotZero(t, portfolio1.ID)
}

func TestUpdateFanStockQuantity(t *testing.T) {
	portfolio := createRandomFanPortfolio(t)

	arg := UpdateFanStockQuantityParams{
		ID:       portfolio.ID,
		Quantity: "65.000000",
	}

	err := testQueries.UpdateFanStockQuantity(context.Background(), arg)
	require.NoError(t, err)
}

func TestDeleteStockFromFanPortfolio(t *testing.T) {
	portfolio := createRandomFanPortfolio(t)
	err := testQueries.DeleteStockFromFanPortfolio(context.Background(), portfolio.StockID)
	require.NoError(t, err)

	portfolio1, err := testQueries.GetFanPortfolio(context.Background(), portfolio.ID)
	require.Empty(t, portfolio1.StockID)
}
