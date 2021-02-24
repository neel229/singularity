package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	db "github.com/neel229/singularity/pkg/db/sqlc"
)

// CreateCurrencyRateRequest holds the json
// data for creating a currency rate record
type CreateCurrencyRateRequest struct {
	CurrencyID     int32  `json:"currency_id"`
	BaseCurrencyID int32  `json:"base_currency_id"`
	Rate           string `json:"rate"`
}

// CreateCurrencyRate creates a new
// currency rate record
func (s *Server) CreateCurrencyRate(ctx context.Context) http.HandlerFunc {
	req := new(CreateCurrencyRateRequest)
	return func(rw http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&req)
		arg := db.CreateCurrencyRateParams{
			CurrencyID:     req.CurrencyID,
			BaseCurrencyID: req.BaseCurrencyID,
			Rate:           req.Rate,
		}
		currencyRate, err := s.store.CreateCurrencyRate(ctx, arg)
		if err != nil {
			http.Error(rw, "error writing data to db", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(rw).Encode(currencyRate)
	}
}

// GetCurrencyRate fetches a currency rate
// based on the id provided
func (s *Server) GetCurrencyRate(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)
		currencyRate, err := s.store.GetCurrencyRate(ctx, id)
		if err != nil {
			http.Error(rw, "error fetching currency rate, check currency rate id", http.StatusBadRequest)
			return
		}
		json.NewEncoder(rw).Encode(currencyRate)
	}
}

type updateCurrencyRateRequest struct {
	Rate string `json:"rate"`
}

// UpdateCurrencyRate updates the currency
// rate.
func (s *Server) UpdateCurrencyRate(ctx context.Context) http.HandlerFunc {
	req := new(updateCurrencyRateRequest)
	return func(rw http.ResponseWriter, r *http.Request) {
		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)
		json.NewDecoder(r.Body).Decode(&req)
		arg := db.UpdateCurrencyRateParams{
			ID:   id,
			Rate: req.Rate,
		}
		if err := s.store.UpdateCurrencyRate(ctx, arg); err != nil {
			http.Error(rw, "error updating currency rate", http.StatusBadRequest)
			return
		}
	}
}
