package api

import (
	"context"
	"encoding/json"
	"net/http"

	db "github.com/neel229/singularity/pkg/db/sqlc"
)

// Logic for creating a creator account
// and corresponding stock
type createCreatorRequest struct {
	FirstName           string `json:"first_name" validate:"required"`
	LastName            string `json:"last_name" validate:"required"`
	UserName            string `json:"user_name" validate:"required"`
	Email               string `json:"email" validate:"required, email"`
	Password            string `json:"password" validate:"required"`
	PreferredCurrencyID int32  `json:"preferred_currency_id" validate:"required"`
	Ticker              string `json:"ticker" validate:"required"`
	Details             string `json:"details" validate:"required"`
}

// CreateCreator executes the StockCreationTX
// and creates a new creator and stock entry
func (s *Server) CreateCreator(ctx context.Context) http.HandlerFunc {
	req := new(createCreatorRequest)
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&req)
		arg := db.StockCreationTxParams{
			FirstName:           req.FirstName,
			LastName:            req.LastName,
			UserName:            req.UserName,
			Email:               req.Email,
			Password:            req.Password,
			PreferredCurrencyID: req.PreferredCurrencyID,
			Ticker:              req.Ticker,
			Details:             req.Details,
		}

		results, err := s.store.StockCreationTx(ctx, arg)
		if err != nil {
			http.Error(w, "error writing data to db", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(results)
	}
}
