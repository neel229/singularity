package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStockCreationTx(t *testing.T) {
	store := NewStore(testDB)

	arg := StockCreationTxParams{
		FirstName:           "Logan",
		LastName:            "Paul",
		UserName:            "loganpaul69",
		Email:               "loganpaul@outlook.com",
		Password:            "klsfjlksqojwo",
		PreferredCurrencyID: 1,
		Ticker:              "LPT",
		Details:             "Stock of Logan Paul",
	}
	result, err := store.StockCreationTx(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	// check Creator account
	creator := result.Creator
	require.Equal(t, arg.FirstName, creator.FirstName)
	require.Equal(t, arg.LastName, creator.LastName)
	require.Equal(t, arg.UserName, creator.UserName)
	require.Equal(t, arg.Email, creator.Email)
	require.Equal(t, arg.Password, creator.Password)
	require.Equal(t, arg.PreferredCurrencyID, creator.PreferredCurrencyID)
	require.NotZero(t, creator.ID)

	_, err = store.GetCreator(context.Background(), creator.ID)
	require.NoError(t, err)

	// check Stock
	stock := result.Stock
	require.Equal(t, arg.Ticker, stock.Ticker)
	require.Equal(t, arg.Details, stock.Details)
	require.NotZero(t, stock.ID)

	_, err = store.GetStock(context.Background(), stock.ID)
	require.NoError(t, err)

	// check Creator & Stock mapping
	creatorStock := result.CreatorStock
	require.Equal(t, creator.ID, creatorStock.CreatorID)
	require.Equal(t, stock.ID, creatorStock.StockID)
	require.NotZero(t, creatorStock.ID)

	_, err = store.GetCreatorStock(context.Background(), creatorStock.ID)
	require.NoError(t, err)

	// check CreatorPortfolio
	portfolio := result.CPortfolio
	require.Equal(t, creator.ID, portfolio.CreatorID)
	require.Equal(t, stock.ID, portfolio.StockID)
	require.Equal(t, "10000000.000000", portfolio.Quantity)
	require.NotZero(t, portfolio.ID)

	_, err = store.GetPortfolioByCreatorID(context.Background(), portfolio.CreatorID)
	require.NoError(t, err)
}

func TestVTrade(t *testing.T) {
	store := NewStore(testDB)

	arg := VTradeTxParams{
		StockID:       1,
		CreatorID:     1,
		FanID:         1,
		Quantity:      "100.000000",
		UnitPrice:     "100.420000",
		Details:       "safjla",
		VirginOfferID: 1,
	}

	result, err := store.VTradeTx(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	// check Virgin Trade
	vTrade := result.VTrade
	require.Equal(t, arg.CreatorID, vTrade.CreatorID)
	require.Equal(t, arg.StockID, vTrade.StockID)
	require.Equal(t, arg.FanID, vTrade.FanID)
	require.Equal(t, arg.Quantity, vTrade.Quantity)
	require.Equal(t, arg.UnitPrice, vTrade.UnitPrice)
	require.Equal(t, arg.Details, vTrade.Details)
	require.Equal(t, arg.VirginOfferID, vTrade.VirginOfferID)
}
