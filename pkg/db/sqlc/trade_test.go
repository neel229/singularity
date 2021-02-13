package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomTrades(t *testing.T) Trade {
	arg := CreateTradeParams{
		StockID:   1,
		BuyerID:   1,
		SellerID:  2,
		Quantity:  "1.000000",
		UnitPrice: "69.420000",
		Details:   "sjaflajoqr",
		OfferID:   1,
	}
	trade, err := testQueries.CreateTrade(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, trade)

	require.Equal(t, arg.StockID, trade.StockID)
	require.Equal(t, arg.BuyerID, trade.BuyerID)
	require.Equal(t, arg.SellerID, trade.SellerID)
	require.Equal(t, arg.Quantity, trade.Quantity)
	require.Equal(t, arg.UnitPrice, trade.UnitPrice)
	require.Equal(t, arg.Details, trade.Details)
	require.Equal(t, arg.OfferID, trade.OfferID)

	require.NotZero(t, trade.ID)

	return trade
}

func TestCreateTrade(t *testing.T) {
	createRandomTrades(t)
}

func TestGetTrade(t *testing.T) {
	trade := createRandomTrades(t)
	trade1, err := testQueries.GetTrade(context.Background(), trade.ID)
	require.NoError(t, err)
	require.NotEmpty(t, trade1)

	require.Equal(t, trade.StockID, trade1.StockID)
	require.Equal(t, trade.BuyerID, trade1.BuyerID)
	require.Equal(t, trade.SellerID, trade1.SellerID)
	require.Equal(t, trade.Quantity, trade1.Quantity)
	require.Equal(t, trade.UnitPrice, trade1.UnitPrice)
	require.Equal(t, trade.Details, trade1.Details)
	require.Equal(t, trade.OfferID, trade1.OfferID)

	require.NotZero(t, trade1.ID)
}

func TestGetTradesByBuyer(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTrades(t)
	}
	arg := GetTradesByBuyerParams{
		BuyerID: 1,
		Limit:   5,
		Offset:  5,
	}
	trades, err := testQueries.GetTradesByBuyer(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, trades, 5)

	for _, trade := range trades {
		require.NotEmpty(t, trade)
	}
}

func TestGetTradesBySeller(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTrades(t)
	}
	arg := GetTradesBySellerParams{
		SellerID: 2,
		Limit:    5,
		Offset:   5,
	}
	trades, err := testQueries.GetTradesBySeller(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, trades, 5)

	for _, trade := range trades {
		require.NotEmpty(t, trade)
	}
}
