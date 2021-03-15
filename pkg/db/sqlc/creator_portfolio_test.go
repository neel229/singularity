package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomCreatorPortfolio(t *testing.T) CreatorPortfolio {
	arg := CreateCreatorPortfolioParams{
		CreatorID: 1,
		StockID:   1,
		Quantity:  "10000000.000000",
	}
	portfolio, err := testQueries.CreateCreatorPortfolio(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, portfolio)

	require.Equal(t, arg.CreatorID, portfolio.CreatorID)
	require.Equal(t, arg.StockID, portfolio.StockID)
	require.Equal(t, arg.Quantity, portfolio.Quantity)

	require.NotZero(t, portfolio.ID)

	return portfolio
}

func TestCreateCreatorPortfolio(t *testing.T) {
	createRandomCreatorPortfolio(t)
}

func TestGetPortfolioByCreatorID(t *testing.T) {
	portfolio := createRandomCreatorPortfolio(t)
	portfolio1, err := testQueries.GetPortfolioByCreatorID(context.Background(), portfolio.CreatorID)
	require.NoError(t, err)
	require.NotEmpty(t, portfolio1)

	require.Equal(t, portfolio.CreatorID, portfolio1.CreatorID)
	require.Equal(t, portfolio.StockID, portfolio1.StockID)
	require.Equal(t, portfolio.Quantity, portfolio1.Quantity)

	require.NotZero(t, portfolio1.ID)
}

func TestUpdateCreatorStockQuantity(t *testing.T) {
	portfolio := createRandomCreatorPortfolio(t)

	arg := UpdateCreatorStockQuantityParams{
		CreatorID: portfolio.CreatorID,
		StockID:   portfolio.StockID,
		Quantity:  "65.000000",
	}

	err := testQueries.UpdateCreatorStockQuantity(context.Background(), arg)
	require.NoError(t, err)
}

func TestDeleteStockFromCreatorPortfolio(t *testing.T) {
	portfolio := createRandomCreatorPortfolio(t)
	arg := DeleteStockFromCreatorPortfolioParams{
		CreatorID: portfolio.CreatorID,
		StockID:   portfolio.StockID,
	}
	err := testQueries.DeleteStockFromCreatorPortfolio(context.Background(), arg)
	require.NoError(t, err)

	portfolio1, err := testQueries.GetPortfolioByCreatorID(context.Background(), portfolio.CreatorID)
	require.NoError(t, err)
	require.Empty(t, portfolio1.StockID)
}
