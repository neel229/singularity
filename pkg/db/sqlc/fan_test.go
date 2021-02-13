package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomFan(t *testing.T) Fan {
	arg := CreateFanParams{
		FirstName:           "Neel",
		LastName:            "Modi",
		UserName:            "neel229",
		Email:               "djfksdjfk",
		Password:            "dfjskfjsd",
		PreferredCurrencyID: 3,
	}
	fan, err := testQueries.CreateFan(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, fan)

	require.Equal(t, arg.FirstName, fan.FirstName)
	require.Equal(t, arg.LastName, fan.LastName)
	require.Equal(t, arg.UserName, fan.UserName)
	require.Equal(t, arg.Email, fan.Email)
	require.Equal(t, arg.Password, fan.Password)
	require.Equal(t, arg.PreferredCurrencyID, fan.PreferredCurrencyID)

	require.NotZero(t, fan.ID)
	require.NotZero(t, fan.TimeConfirmed)
	require.NotZero(t, fan.TimeRegistered)

	return fan
}

func TestCreateFan(t *testing.T) {
	createRandomFan(t)
}

func TestGetFan(t *testing.T) {
	fan := createRandomFan(t)
	fan1, err := testQueries.GetFan(context.Background(), fan.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fan1)

	require.Equal(t, fan.FirstName, fan1.FirstName)
	require.Equal(t, fan.LastName, fan1.LastName)
	require.Equal(t, fan.UserName, fan1.UserName)
	require.Equal(t, fan.Email, fan1.Email)
	require.Equal(t, fan.Password, fan1.Password)
	require.Equal(t, fan.PreferredCurrencyID, fan1.PreferredCurrencyID)

	require.NotZero(t, fan1.ID)
	require.NotZero(t, fan1.TimeConfirmed)
	require.NotZero(t, fan1.TimeRegistered)
}

func TestUpdateEmail(t *testing.T) {
	fan := createRandomFan(t)
	arg := UpdateEmailParams{
		ID:    fan.ID,
		Email: "qowieuqwoelan",
	}

	err := testQueries.UpdateEmail(context.Background(), arg)
	require.NoError(t, err)
}

func TestUpdatePassword(t *testing.T) {
	fan := createRandomFan(t)
	arg := UpdatePasswordParams{
		ID:       fan.ID,
		Password: "acmvmxvx",
	}

	err := testQueries.UpdatePassword(context.Background(), arg)
	require.NoError(t, err)
}

func TestUpdatePreferredCurrency(t *testing.T) {
	fan := createRandomFan(t)
	arg := UpdatePreferredCurrencyParams{
		ID:                  fan.ID,
		PreferredCurrencyID: 2,
	}

	err := testQueries.UpdatePreferredCurrency(context.Background(), arg)
	require.NoError(t, err)
}
