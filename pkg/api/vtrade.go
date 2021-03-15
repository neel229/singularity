package api

import (
	"context"
	"encoding/json"
	"net/http"

	db "github.com/neel229/forum/pkg/db/sqlc"
)

type vTradeTxRequest struct {
	StockID   int64  `json:"stock_id"`
	CreatorID int64  `json:"creator_id"`
	FanID     int64  `json:"fan_id"`
	Quantity  string `json:"quantity"`
	UnitPrice string `json:"unit_price"`
	Details   string `json:"details"`
}

// VTradeTx is the transaction which takes place
// between a creator and a fan. Creator's tokens are
// transferred from creator's account to fan's account
func (s *Server) VTradeTx(ctx context.Context) http.HandlerFunc {
	req := new(vTradeTxRequest)
	return func(rw http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&req)
		arg := db.VTradeTxParams{
			StockID:   req.StockID,
			CreatorID: req.CreatorID,
			FanID:     req.FanID,
			Quantity:  req.Quantity,
			UnitPrice: req.UnitPrice,
			Details:   req.Details,
		}
		result, err := s.store.VTradeTx(ctx, arg)
		if err != nil {
			http.Error(rw, "there was an error executing transaction", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(rw).Encode(result)
	}
}
