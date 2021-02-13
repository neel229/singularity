package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomReport(t *testing.T) Report {
	arg := CreateReportParams{
		TradingDate: time.Date(2021, time.February, time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.Now().UTC().Location()),
		StockID:     1,
		CurrencyID:  1,
		FirstPrice:  "100.000000",
		LastPrice:   "102.000000",
		MinPrice:    "99.000000",
		MaxPrice:    "104.000000",
		AvgPrice:    "101.500000",
		TotalAmount: "1000000000.000000",
		Volume:      "10000000.000000",
	}

	report, err := testQueries.CreateReport(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, report)

	require.Equal(t, arg.StockID, report.StockID)
	require.Equal(t, arg.CurrencyID, report.CurrencyID)
	require.Equal(t, arg.FirstPrice, report.FirstPrice)
	require.Equal(t, arg.LastPrice, report.LastPrice)
	require.Equal(t, arg.MinPrice, report.MinPrice)
	require.Equal(t, arg.MaxPrice, report.MaxPrice)
	require.Equal(t, arg.AvgPrice, report.AvgPrice)
	require.Equal(t, arg.TotalAmount, report.TotalAmount)
	require.Equal(t, arg.Volume, report.Volume)

	return report
}

func TestCreateReport(t *testing.T) {
	createRandomReport(t)
}

func TestGetReport(t *testing.T) {
	report := createRandomReport(t)

	report1, err := testQueries.GetReport(context.Background(), report.ID)
	require.NoError(t, err)
	require.NotEmpty(t, report1)

	require.Equal(t, report.StockID, report.StockID)
	require.Equal(t, report.CurrencyID, report1.CurrencyID)
	require.Equal(t, report.FirstPrice, report1.FirstPrice)
	require.Equal(t, report.LastPrice, report1.LastPrice)
	require.Equal(t, report.MinPrice, report1.MinPrice)
	require.Equal(t, report.MaxPrice, report1.MaxPrice)
	require.Equal(t, report.AvgPrice, report1.AvgPrice)
	require.Equal(t, report.TotalAmount, report1.TotalAmount)
	require.Equal(t, report.Volume, report1.Volume)
}
