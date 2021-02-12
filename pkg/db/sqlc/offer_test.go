package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomOffers(t *testing.T) Offer {
	arg := CreateOfferParams{
		TraderID: 3,
		StockID:  1,
		Quantity: "1.000000",
		Buy:      true,
		Sell:     false,
		Price:    "69.420000",
	}
	offer, err := testQueries.CreateOffer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, offer)

	require.Equal(t, arg.TraderID, offer.TraderID)
	require.Equal(t, arg.StockID, offer.StockID)
	require.Equal(t, arg.Quantity, offer.Quantity)
	require.Equal(t, arg.Buy, offer.Buy)
	require.Equal(t, arg.Sell, offer.Sell)
	require.Equal(t, arg.Price, offer.Price)

	require.NotZero(t, offer.ID)

	return offer
}

func TestCreateOffer(t *testing.T) {
	createRandomOffers(t)
}

func TestGetOffer(t *testing.T) {
	offer := createRandomOffers(t)
	offer1, err := testQueries.GetOffer(context.Background(), offer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, offer1)

	require.Equal(t, offer.TraderID, offer1.TraderID)
	require.Equal(t, offer.StockID, offer1.StockID)
	require.Equal(t, offer.Quantity, offer1.Quantity)
	require.Equal(t, offer.Buy, offer1.Buy)
	require.Equal(t, offer.Sell, offer1.Sell)
	require.Equal(t, offer.Price, offer1.Price)

	require.NotZero(t, offer1.ID)
}

func TestListOffersByTrader(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomOffers(t)
	}

	arg := ListOffersParams{
		TraderID: 1,
		Limit:    5,
		Offset:   5,
	}

	offers, err := testQueries.ListOffers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, offers, 5)

	for _, offer := range offers {
		require.NotEmpty(t, offer)
	}
}
