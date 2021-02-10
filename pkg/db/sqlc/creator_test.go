package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateCreator(t *testing.T) {
	arg := CreateCreatorParams{
		FirstName:           "David",
		LastName:            "Dobrik",
		UserName:            "daviddobrik",
		Email:               "djfksdjfk",
		Password:            "dfjskfjsd",
		PreferredCurrencyID: 2,
		CreatorStockID:      sql.NullInt64{Int64: 1, Valid: true},
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
	require.Equal(t, arg.CreatorStockID, creator.CreatorStockID)
	require.Equal(t, arg.VirginTokensLeft, creator.VirginTokensLeft)

	require.NotZero(t, creator.TimeConfirmed)
	require.NotZero(t, creator.TimeRegistered)
}
