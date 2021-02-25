package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	db "github.com/neel229/singularity/pkg/db/sqlc"
)

type createFanRequest struct {
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	UserName            string `json:"user_name"`
	Password            string `json:"password"`
	Email               string `json:"email"`
	PreferredCurrencyID int64  `json:"preferred_currency_id"`
}

// CreateFan creates a new fan account
func (s *Server) CreateFan(ctx context.Context) http.HandlerFunc {
	req := new(createFanRequest)
	return func(rw http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&req)
		arg := db.CreateFanParams{
			FirstName:           req.FirstName,
			LastName:            req.LastName,
			UserName:            req.UserName,
			Password:            req.Password,
			Email:               req.Email,
			PreferredCurrencyID: req.PreferredCurrencyID,
		}
		fan, err := s.store.CreateFan(ctx, arg)
		if err != nil {
			http.Error(rw, "error creating fan", http.StatusBadRequest)
			return
		}
		json.NewEncoder(rw).Encode(fan)
	}
}

// GetFan fetches a fan's details from the provided id
func (s *Server) GetFan(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)
		fan, err := s.store.GetFan(ctx, id)
		if err != nil {
			http.Error(rw, "check the fan id", http.StatusBadRequest)
			return
		}
		json.NewEncoder(rw).Encode(fan)
	}
}

type updateEmailRequest struct {
	Email string `json:"email"`
}

// UpdateFanEmail updates the email id
// of a fan account based on the id provided
func (s *Server) UpdateFanEmail(ctx context.Context) http.HandlerFunc {
	req := new(updateEmailRequest)
	return func(rw http.ResponseWriter, r *http.Request) {
		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)
		json.NewDecoder(r.Body).Decode(req)
		arg := db.UpdateEmailParams{
			ID:    id,
			Email: req.Email,
		}
		if err := s.store.UpdateEmail(ctx, arg); err != nil {
			http.Error(rw, "error updating email, check back later!", http.StatusInternalServerError)
			return
		}
	}
}

type updatePasswordRequest struct {
	Password string `json:"password"`
}

// UpdateFanPassword updates the password
// of a fan account based on the id provided
func (s *Server) UpdateFanPassword(ctx context.Context) http.HandlerFunc {
	req := new(updatePasswordRequest)
	return func(rw http.ResponseWriter, r *http.Request) {
		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)
		json.NewDecoder(r.Body).Decode(req)
		arg := db.UpdatePasswordParams{
			ID:       id,
			Password: req.Password,
		}
		if err := s.store.UpdatePassword(ctx, arg); err != nil {
			http.Error(rw, "error updating password, check back later!", http.StatusInternalServerError)
			return
		}
	}
}

type updatePreferredCurrencyRequest struct {
	PreferredCurrencyID int64 `json:"preferred_currency_id"`
}

// UpdateFanPreferredCurrency updates the preferred
// currency of a fan account based on the id provided
func (s *Server) UpdateFanPreferredCurrency(ctx context.Context) http.HandlerFunc {
	req := new(updatePreferredCurrencyRequest)
	return func(rw http.ResponseWriter, r *http.Request) {
		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)
		json.NewDecoder(r.Body).Decode(req)
		arg := db.UpdatePreferredCurrencyParams{
			ID:                  id,
			PreferredCurrencyID: req.PreferredCurrencyID,
		}
		if err := s.store.UpdatePreferredCurrency(ctx, arg); err != nil {
			http.Error(rw, "error updating preferred currency, check back later!", http.StatusInternalServerError)
			return
		}
	}
}
