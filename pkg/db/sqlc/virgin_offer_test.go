package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomVirginOffers(t *testing.T) VirginOffer {
	arg := CreateVirginOfferParams{
		CreatorID: 1,
		StockID:   1,
		Quantity:  "1.000000",
		Price:     "69.420000",
	}
	offer, err := testQueries.CreateVirginOffer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, offer)

	require.Equal(t, arg.CreatorID, offer.CreatorID)
	require.Equal(t, arg.StockID, offer.StockID)
	require.Equal(t, arg.Quantity, offer.Quantity)
	require.Equal(t, arg.Price, offer.Price)

	require.NotZero(t, offer.ID)

	return offer
}

func TestCreateVirginOffer(t *testing.T) {
	createRandomVirginOffers(t)
}

func TestGetVirginOffer(t *testing.T) {
	offer := createRandomVirginOffers(t)
	offer1, err := testQueries.GetVirginOffer(context.Background(), offer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, offer1)

	require.Equal(t, offer.CreatorID, offer1.CreatorID)
	require.Equal(t, offer.StockID, offer1.StockID)
	require.Equal(t, offer.Quantity, offer1.Quantity)
	require.Equal(t, offer.Price, offer1.Price)

	require.NotZero(t, offer1.ID)
}
func TestGetVirginOfferByCreator(t *testing.T) {
	offer := createRandomVirginOffers(t)
	offer1, err := testQueries.GetVirginOfferByCreator(context.Background(), offer.CreatorID)

	require.NoError(t, err)
	require.NotEmpty(t, offer1)

	require.Equal(t, offer.CreatorID, offer1.CreatorID)
	require.Equal(t, offer.StockID, offer1.StockID)
	require.Equal(t, offer.Quantity, offer1.Quantity)
	require.Equal(t, offer.Price, offer1.Price)

	require.NotZero(t, offer1.ID)
}

func TestListOffersByCreator(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomVirginOffers(t)
	}

	arg := ListVirginOffersParams{
		Limit:  5,
		Offset: 5,
	}

	offers, err := testQueries.ListVirginOffers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, offers, 5)

	for _, offer := range offers {
		require.NotEmpty(t, offer)
	}
}
