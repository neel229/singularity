package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	db "github.com/neel229/forum/pkg/db/sqlc"
)

type createCurrencyRequest struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	IsBase bool   `json:"is_base"`
}

// CreateCurrency adds a currency
// to the list of currencies available
// for trading on the platform
func (s *Server) CreateCurrency(ctx context.Context) http.HandlerFunc {
	req := new(createCurrencyRequest)
	return func(rw http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&req)
		arg := db.CreateCurrencyParams{
			Code:   req.Code,
			Name:   req.Name,
			IsBase: req.IsBase,
		}

		currency, err := s.store.CreateCurrency(ctx, arg)
		if err != nil {
			http.Error(rw, "error writing data to db", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(rw).Encode(currency)
	}
}

// GetCurrency fetches the currency
// based on the id supplied
func (s *Server) GetCurrency(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Access-Control-Allow-Origin", "*")
		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)
		currency, err := s.store.GetCurrency(ctx, id)
		if err != nil {
			http.Error(rw, "invalid currency", http.StatusBadRequest)
			return
		}
		json.NewEncoder(rw).Encode(&currency)
	}
}

type listCurrenciesRequest struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

// ListCurrencies lists all the currencies
// tradable on the platform
func (s *Server) ListCurrencies(ctx context.Context) http.HandlerFunc {
	req := new(listCurrenciesRequest)
	return func(rw http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&req)
		arg := db.ListCurrenciesParams{
			Limit:  req.Limit,
			Offset: req.Offset,
		}
		currencies, err := s.store.ListCurrencies(ctx, arg)
		if err != nil {
			http.Error(rw, "problem fetching the list of currenices", http.StatusInternalServerError)
		}
		json.NewEncoder(rw).Encode(currencies)
	}
}

type updateCurrencyRequest struct {
	ID     int64 `json:"id"`
	IsBase bool  `json:"is_base"`
}

// UpdateCurrency updates the isBase field
// of the given currency
func (s *Server) UpdateCurrency(ctx context.Context) http.HandlerFunc {
	req := new(updateCurrencyRequest)
	return func(rw http.ResponseWriter, r *http.Request) {
		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)
		json.NewDecoder(r.Body).Decode(&req)
		arg := db.UpdateCurrencyParams{
			ID:     id,
			IsBase: req.IsBase,
		}
		if err := s.store.UpdateCurrency(ctx, arg); err != nil {
			http.Error(rw, "error updating the given currency", http.StatusBadRequest)
			return
		}
	}
}

// DeleteCurrency deletes the currency
// with the provided id from.
func (s *Server) DeleteCurrency(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)
		if err := s.store.DeleteCurrency(ctx, id); err != nil {
			http.Error(rw, "error deleting the currency, check the currency id", http.StatusBadRequest)
		}
	}
}
