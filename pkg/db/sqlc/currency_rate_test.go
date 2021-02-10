package db

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/neel229/singularity/pkg/util"
	"github.com/stretchr/testify/require"
)

func createRandomCurrencyRate(t *testing.T) CurrencyRate {
	// arg := CreateCurrencyRateParams{
	// 	CurrencyID:     int32(util.RandomInt(1, 5)),
	// 	BaseCurrencyID: int32(util.RandomInt(2, 7)),
	// 	Rate:           strconv.Itoa(int(util.RandomInt(20, 100))) + ".000000",
	// }

	arg := CreateCurrencyRateParams{
		CurrencyID: 1,
		BaseCurrencyID: 2,
		Rate: "20.000000",
	}
	currencyRate, err := testQueries.CreateCurrencyRate(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, currencyRate)

	require.Equal(t, arg.CurrencyID, currencyRate.CurrencyID)
	require.Equal(t, arg.BaseCurrencyID, currencyRate.BaseCurrencyID)
	require.Equal(t, arg.Rate, currencyRate.Rate)

	require.NotZero(t, currencyRate.ID)
	require.NotZero(t, currencyRate.Ts)

	return currencyRate
}

func TestCreateCurrencyRate(t *testing.T) {
	createRandomCurrencyRate(t)
}

func TestGetCurrencyRate(t *testing.T) {
	currencyRate1 := createRandomCurrencyRate(t)
	currencyRate2, err := testQueries.GetCurrencyRate(context.Background(), int64(currencyRate1.ID))
	require.NoError(t, err)
	require.NotEmpty(t, currencyRate2)

	require.Equal(t, currencyRate1.CurrencyID, currencyRate2.CurrencyID)
	require.Equal(t, currencyRate1.BaseCurrencyID, currencyRate2.BaseCurrencyID)
	require.Equal(t, currencyRate1.Rate, currencyRate2.Rate)

	require.WithinDuration(t, currencyRate1.Ts, currencyRate2.Ts, time.Second)
}

func TestUpdateCurrencyRate(t *testing.T) {
	currencyRate1 := createRandomCurrencyRate(t)
	arg := UpdateCurrencyRateParams{
		ID:   currencyRate1.ID,
		Rate: strconv.Itoa(int(util.RandomInt(20, 100))) + ".000000",
	}
	err := testQueries.UpdateCurrencyRate(context.Background(), arg)
	require.NoError(t, err)
}
