package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomPortfolios(t *testing.T) Portfolio {
	arg := CreatePortfolioParams{
		TraderID: 1,
		StockID:  1,
		Quantity: "50.000000",
	}
	portfolio, err := testQueries.CreatePortfolio(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, portfolio)

	require.Equal(t, arg.TraderID, portfolio.TraderID)
	require.Equal(t, arg.StockID, portfolio.StockID)
	require.Equal(t, arg.Quantity, portfolio.Quantity)

	require.NotZero(t, portfolio.ID)

	return portfolio
}

func TestCreatePortfolio(t *testing.T) {
	createRandomPortfolios(t)
}

func TestGetPortfolio(t *testing.T) {
	portfolio := createRandomPortfolios(t)
	portfolio1, err := testQueries.GetPortfolio(context.Background(), portfolio.ID)
	require.NoError(t, err)
	require.NotEmpty(t, portfolio1)

	require.Equal(t, portfolio.TraderID, portfolio1.TraderID)
	require.Equal(t, portfolio.StockID, portfolio1.StockID)
	require.Equal(t, portfolio.Quantity, portfolio1.Quantity)

	require.NotZero(t, portfolio1.ID)
}

func TestGetPortfolioByTraderID(t *testing.T) {
	portfolio := createRandomPortfolios(t)
	portfolio1, err := testQueries.GetPortfolioByTraderID(context.Background(), portfolio.TraderID)
	require.NoError(t, err)
	require.NotEmpty(t, portfolio1)

	require.Equal(t, portfolio.TraderID, portfolio1.TraderID)
	require.Equal(t, portfolio.StockID, portfolio1.StockID)
	require.Equal(t, portfolio.Quantity, portfolio1.Quantity)

	require.NotZero(t, portfolio1.ID)
}

func TestUpdateStockQuantity(t *testing.T) {
	portfolio := createRandomPortfolios(t)

	arg := UpdateStockQuantityParams{
		ID:       portfolio.ID,
		Quantity: "65.000000",
	}

	err := testQueries.UpdateStockQuantity(context.Background(), arg)
	require.NoError(t, err)
}

func TestDeleteStockFromPortfolio(t *testing.T) {
	portfolio := createRandomPortfolios(t)
	err := testQueries.DeleteStockFromPortfolio(context.Background(), portfolio.StockID)
	require.NoError(t, err)

	portfolio1, err := testQueries.GetPortfolio(context.Background(), portfolio.ID)
	require.Empty(t, portfolio1.StockID)
}
