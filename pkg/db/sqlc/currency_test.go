package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/neel229/singularity/pkg/util"
	"github.com/stretchr/testify/require"
)

func createRandomCurrency(t *testing.T) Currency {
	code := util.RandomCurrencyCode()
	name := util.CurrencyName(code)
	isBase := util.IsBase(code)
	arg := CreateCurrencyParams{
		Code:   code,
		Name:   name,
		IsBase: isBase,
	}
	currency, err := testQueries.CreateCurrency(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, currency)

	require.Equal(t, arg.Code, currency.Code)
	require.Equal(t, arg.Name, currency.Name)
	require.Equal(t, arg.IsBase, currency.IsBase)

	require.NotZero(t, currency.ID)

	return currency
}

func TestCreateCurrency(t *testing.T) {
	createRandomCurrency(t)
}

func TestGetCurrency(t *testing.T) {
	currency := createRandomCurrency(t)

	currency1, err := testQueries.GetCurrency(context.Background(), currency.ID)
	require.NoError(t, err)
	require.NotEmpty(t, currency1)

	require.Equal(t, currency.Code, currency1.Code)
	require.Equal(t, currency.Name, currency1.Name)
	require.Equal(t, currency.IsBase, currency1.IsBase)
	require.Equal(t, currency.ID, currency1.ID)
}

func TestUpdateCurrency(t *testing.T) {
	currency := createRandomCurrency(t)
	arg := UpdateCurrencyParams{
		ID:     currency.ID,
		IsBase: !currency.IsBase,
	}
	err := testQueries.UpdateCurrency(context.Background(), arg)
	require.NoError(t, err)
}

func TestDeleteAccount(t *testing.T) {
	currency := createRandomCurrency(t)
	err := testQueries.DeleteCurrency(context.Background(), currency.ID)
	require.NoError(t, err)
	currency1, err := testQueries.GetCurrency(context.Background(), currency.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, currency1)
}

func TestListCurrencies(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCurrency(t)
	}

	arg := ListCurrenciesParams{
		Limit:  5,
		Offset: 5,
	}

	currencies, err := testQueries.ListCurrencies(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, currencies, 5)

	for _, currency := range currencies {
		require.NotEmpty(t, currency)
	}
}
