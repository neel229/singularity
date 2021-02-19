package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomVirginTrades(t *testing.T) VirginTrade {
	arg := CreateVirginTradeParams{
		StockID:       1,
		FanID:         1,
		CreatorID:     2,
		Quantity:      "1.000000",
		UnitPrice:     "69.420000",
		Details:       "sjaflajoqr",
		VirginOfferID: 1,
	}
	trade, err := testQueries.CreateVirginTrade(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, trade)

	require.Equal(t, arg.StockID, trade.StockID)
	require.Equal(t, arg.FanID, trade.FanID)
	require.Equal(t, arg.CreatorID, trade.CreatorID)
	require.Equal(t, arg.Quantity, trade.Quantity)
	require.Equal(t, arg.UnitPrice, trade.UnitPrice)
	require.Equal(t, arg.Details, trade.Details)
	require.Equal(t, arg.VirginOfferID, trade.VirginOfferID)

	require.NotZero(t, trade.ID)

	return trade
}

func TestCreateVirginTrade(t *testing.T) {
	createRandomVirginTrades(t)
}

func TestGetVirginTrade(t *testing.T) {
	trade := createRandomVirginTrades(t)
	trade1, err := testQueries.GetVirginTrade(context.Background(), trade.ID)
	require.NoError(t, err)
	require.NotEmpty(t, trade1)

	require.Equal(t, trade.StockID, trade1.StockID)
	require.Equal(t, trade.FanID, trade1.FanID)
	require.Equal(t, trade.CreatorID, trade1.CreatorID)
	require.Equal(t, trade.Quantity, trade1.Quantity)
	require.Equal(t, trade.UnitPrice, trade1.UnitPrice)
	require.Equal(t, trade.Details, trade1.Details)
	require.Equal(t, trade.VirginOfferID, trade1.VirginOfferID)

	require.NotZero(t, trade1.ID)
}

func TestListVirginTradesByCreator(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomVirginTrades(t)
	}
	arg := ListVirginTradesByCreatorParams{
		CreatorID: 2,
		Limit:     5,
		Offset:    5,
	}
	trades, err := testQueries.ListVirginTradesByCreator(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, trades, 5)

	for _, trade := range trades {
		require.NotEmpty(t, trade)
	}
}

func TestVirginTradesByFans(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomVirginTrades(t)
	}
	arg := ListVirginTradesByFanParams{
		FanID:  1,
		Limit:  5,
		Offset: 5,
	}
	trades, err := testQueries.ListVirginTradesByFan(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, trades, 5)

	for _, trade := range trades {
		require.NotEmpty(t, trade)
	}
}
