package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomCreator(t *testing.T) Creator {
	arg := CreateCreatorParams{
		FirstName:           "Logan",
		LastName:            "Paul",
		UserName:            "loganpaul",
		Email:               "djfksdjfk",
		Password:            "dfjskfjsd",
		PreferredCurrencyID: 3,
		VirginTokensLeft:    100000,
	}

	creator, err := testQueries.CreateCreator(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, creator)

	require.Equal(t, arg.FirstName, creator.FirstName)
	require.Equal(t, arg.LastName, creator.LastName)
	require.Equal(t, arg.UserName, creator.UserName)
	require.Equal(t, arg.Email, creator.Email)
	require.Equal(t, arg.Password, creator.Password)
	require.Equal(t, arg.PreferredCurrencyID, creator.PreferredCurrencyID)
	require.Equal(t, arg.VirginTokensLeft, creator.VirginTokensLeft)

	require.NotZero(t, creator.TimeConfirmed)
	require.NotZero(t, creator.TimeRegistered)

	return creator
}

func TestCreateCreator(t *testing.T) {
	createRandomCreator(t)
}

func TestGetCreator(t *testing.T) {
	creator := createRandomCreator(t)
	creator1, err := testQueries.GetCreator(context.Background(), creator.ID)
	require.NoError(t, err)
	require.NotEmpty(t, creator1)

	require.Equal(t, creator.FirstName, creator1.FirstName)
	require.Equal(t, creator.LastName, creator1.LastName)
	require.Equal(t, creator.UserName, creator1.UserName)
	require.Equal(t, creator.Email, creator1.Email)
	require.Equal(t, creator.Password, creator1.Password)
	require.Equal(t, creator.PreferredCurrencyID, creator1.PreferredCurrencyID)
	require.Equal(t, creator.VirginTokensLeft, creator1.VirginTokensLeft)

	require.NotZero(t, creator1.TimeConfirmed)
	require.NotZero(t, creator1.TimeRegistered)
}

func TestGetVirginTokensLeft(t *testing.T) {
	creator := createRandomCreator(t)
	tokens, err := testQueries.GetVirginTokensLeft(context.Background(), creator.ID)
	require.NoError(t, err)
	require.Equal(t, creator.VirginTokensLeft, tokens)
}

func TestListCreators(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCreator(t)
	}

	arg := ListCreatorsParams{
		Limit:  5,
		Offset: 5,
	}

	creators, err := testQueries.ListCreators(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, creators, 5)

	for _, creator := range creators {
		require.NotEmpty(t, creator)
	}
}

func TestUpdateCreatorEmail(t *testing.T) {
	creator := createRandomCreator(t)
	arg := UpdateCreatorEmailParams{
		ID:    creator.ID,
		Email: "qowieuqwoelan",
	}

	err := testQueries.UpdateCreatorEmail(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, creator.Email)
}

func TestUpdateCreatorPassword(t *testing.T) {
	creator := createRandomCreator(t)
	arg := UpdateCreatorPasswordParams{
		ID:       creator.ID,
		Password: "acmvmxvx",
	}

	err := testQueries.UpdateCreatorPassword(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, creator.Password)
}

func TestUpdateCreatorPreferredCurrency(t *testing.T) {
	creator := createRandomCreator(t)
	arg := UpdateCreatorPreferredCurrencyParams{
		ID:                  creator.ID,
		PreferredCurrencyID: 2,
	}

	err := testQueries.UpdateCreatorPreferredCurrency(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, creator.PreferredCurrencyID)
}

func TestUpdateVirginTokensLeft(t *testing.T) {
	creator := createRandomCreator(t)
	arg := UpdateVirginTokensLeftParams{
		ID:               creator.ID,
		VirginTokensLeft: 100,
	}

	err := testQueries.UpdateVirginTokensLeft(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, creator.PreferredCurrencyID)
}
