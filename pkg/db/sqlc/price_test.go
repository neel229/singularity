package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomPrices(t *testing.T) Price {
	arg := CreatePriceParams{
		StockID:    1,
		CurrencyID: 1,
		Buy:        "100.000000",
		Sell:       "110.000000",
	}
	price, err := testQueries.CreatePrice(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, price)

	require.Equal(t, arg.StockID, price.StockID)
	require.Equal(t, arg.CurrencyID, price.CurrencyID)
	require.Equal(t, arg.Buy, price.Buy)
	require.Equal(t, arg.Sell, price.Sell)

	require.NotZero(t, price.ID)
	require.NotZero(t, price.Ts)

	return price
}

func TestCreatePrice(t *testing.T) {
	createRandomPrices(t)
}

func TestGetPrice(t *testing.T) {
	price := createRandomPrices(t)
	price1, err := testQueries.GetPrice(context.Background(), price.ID)
	require.NoError(t, err)
	require.NotEmpty(t, price)

	require.Equal(t, price.StockID, price1.StockID)
	require.Equal(t, price.CurrencyID, price1.CurrencyID)
	require.Equal(t, price.Buy, price1.Buy)
	require.Equal(t, price.Sell, price1.Sell)

	require.NotZero(t, price1.ID)
	require.WithinDuration(t, price.Ts, price1.Ts, time.Second)
}

func TestUpdateBuyPrice(t *testing.T) {
	price := createRandomPrices(t)
	arg := UpdateBuyingPriceParams{
		ID:  price.ID,
		Buy: "200.000000",
	}
	err := testQueries.UpdateBuyingPrice(context.Background(), arg)
	require.NoError(t, err)
}

func TestUpdateSellPrice(t *testing.T) {
	price := createRandomPrices(t)
	arg := UpdateSellingPriceParams{
		ID:   price.ID,
		Sell: "230.000000",
	}
	err := testQueries.UpdateSellingPrice(context.Background(), arg)
	require.NoError(t, err)
}

func TestDeletePrice(t *testing.T) {
	price := createRandomPrices(t)
	err := testQueries.DeletePrice(context.Background(), price.ID)
	require.NoError(t, err)
	price1, err := testQueries.GetPrice(context.Background(), price.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, price1)
}
