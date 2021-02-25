package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
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
	PreferredCurrencyID int64  `json:"preferred_currency_id" validate:"required"`
	Ticker              string `json:"ticker" validate:"required"`
	Details             string `json:"details" validate:"required"`
}

// CreateCreatorAndStock executes the StockCreationTX
// and creates a new creator and stock entry
func (s *Server) CreateCreatorAndStock(ctx context.Context) http.HandlerFunc {
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

// CRUD operations on creator

// GetCreator fetches the creator id
// based on the id provided
func (s *Server) GetCreator(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)
		creator, err := s.store.GetCreator(ctx, id)
		if err != nil {
			http.Error(rw, "error fetching creator, check id", http.StatusBadRequest)
			return
		}
		json.NewEncoder(rw).Encode(creator)
	}
}

// GetVirginTokensLeft fetches the amount of
// virgin tokens left with the creator
func (s *Server) GetVirginTokensLeft(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)
		amount, err := s.store.GetVirginTokensLeft(ctx, id)
		if err != nil {
			http.Error(rw, "error fetching tokens left, check id", http.StatusBadRequest)
			return
		}
		json.NewEncoder(rw).Encode(amount)
	}
}

type listCreatorsRequest struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

// ListCreators returns the list of creators
// present on the platform
func (s *Server) ListCreators(ctx context.Context) http.HandlerFunc {
	req := new(listCreatorsRequest)
	return func(rw http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&req)
		arg := db.ListCreatorsParams{
			Limit:  req.Limit,
			Offset: req.Offset,
		}
		creators, err := s.store.ListCreators(ctx, arg)
		if err != nil {
			http.Error(rw, "error returning list of creators", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(rw).Encode(creators)
	}
}

type updateCreatorEmailRequest struct {
	Email string `json:"email"`
}

// UpdateCreatorEmail updates the email id
// of a fan account based on the id provided
func (s *Server) UpdateCreatorEmail(ctx context.Context) http.HandlerFunc {
	req := new(updateCreatorEmailRequest)
	return func(rw http.ResponseWriter, r *http.Request) {
		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)
		json.NewDecoder(r.Body).Decode(req)
		arg := db.UpdateCreatorEmailParams{
			ID:    id,
			Email: req.Email,
		}
		if err := s.store.UpdateCreatorEmail(ctx, arg); err != nil {
			http.Error(rw, "error updating email, check back later!", http.StatusInternalServerError)
			return
		}
	}
}

type updateCreatorPasswordRequest struct {
	Password string `json:"password"`
}

// UpdateCreatorPassword updates the password
// of a fan account based on the id provided
func (s *Server) UpdateCreatorPassword(ctx context.Context) http.HandlerFunc {
	req := new(updateCreatorPasswordRequest)
	return func(rw http.ResponseWriter, r *http.Request) {
		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)
		json.NewDecoder(r.Body).Decode(req)
		arg := db.UpdateCreatorPasswordParams{
			ID:       id,
			Password: req.Password,
		}
		if err := s.store.UpdateCreatorPassword(ctx, arg); err != nil {
			http.Error(rw, "error updating password, check back later!", http.StatusInternalServerError)
			return
		}
	}
}

type updateCreatorPreferredCurrencyRequest struct {
	PreferredCurrencyID int64 `json:"preferred_currency_id"`
}

// UpdateCreatorPreferredCurrency updates the preferred
// currency of a fan account based on the id provided
func (s *Server) UpdateCreatorPreferredCurrency(ctx context.Context) http.HandlerFunc {
	req := new(updateCreatorPreferredCurrencyRequest)
	return func(rw http.ResponseWriter, r *http.Request) {
		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)
		json.NewDecoder(r.Body).Decode(req)
		arg := db.UpdateCreatorPreferredCurrencyParams{
			ID:                  id,
			PreferredCurrencyID: req.PreferredCurrencyID,
		}
		if err := s.store.UpdateCreatorPreferredCurrency(ctx, arg); err != nil {
			http.Error(rw, "error updating preferred currency, check back later!", http.StatusInternalServerError)
			return
		}
	}
}

type updateCreatorVTRequest struct {
	VirginTokensLeft int32 `json:"virgin_tokens_left"`
}

// UpdateCreatorVTLeft  updates the virgin
// tokens left column of creator table
func (s *Server) UpdateCreatorVTLeft(ctx context.Context) http.HandlerFunc {
	req := new(updateCreatorVTRequest)
	return func(rw http.ResponseWriter, r *http.Request) {
		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)
		json.NewDecoder(r.Body).Decode(&req)
		arg := db.UpdateVirginTokensLeftParams{
			ID:               id,
			VirginTokensLeft: req.VirginTokensLeft,
		}
		if err := s.store.UpdateVirginTokensLeft(ctx, arg); err != nil {
			http.Error(rw, "error updating virign tokens left", http.StatusInternalServerError)
			return
		}
	}
}

// CRUD operations on stocks

// GetStock fetches the stock respect to
// provided id
func (s *Server) GetStock(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)
		stock, err := s.store.GetStock(ctx, id)
		if err != nil {
			http.Error(rw, "error fetching the stock", http.StatusBadRequest)
			return
		}
		json.NewEncoder(rw).Encode(stock)
	}
}

type listStocksRequest struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

// ListStocks returns a list of stocks
// available for trading on the platform
func (s *Server) ListStocks(ctx context.Context) http.HandlerFunc {
	req := new(listStocksRequest)
	return func(rw http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&req)
		arg := db.ListStocksParams{
			Limit:  req.Limit,
			Offset: req.Offset,
		}
		stocks, err := s.store.ListStocks(ctx, arg)
		if err != nil {
			http.Error(rw, "error returning list of stocks", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(rw).Encode(stocks)
	}
}

type updateStockDetailsRequest struct {
	Details string `json:"details"`
}

// UpdateStockDetails updates the details of the
// creator's stock
func (s *Server) UpdateStockDetails(ctx context.Context) http.HandlerFunc {
	req := new(updateStockDetailsRequest)
	return func(rw http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&req)
		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)
		arg := db.UpdateStockParams{
			ID:      id,
			Details: req.Details,
		}
		if err := s.store.UpdateStock(ctx, arg); err != nil {
			http.Error(rw, "error updating the details of the stock", http.StatusBadRequest)
			return
		}
	}
}

// CRUD operations for creator_stock

// GetCreatorStock fetches the stock
// from the provided creator_stock id
func (s *Server) GetCreatorStock(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)

		creatorStock, err := s.store.GetCreatorStock(ctx, id)
		if err != nil {
			http.Error(rw, "error fetching the creator stock", http.StatusBadRequest)
			return
		}
		json.NewEncoder(rw).Encode(creatorStock)
	}
}

type listCreatorStockRequest struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

// ListCreatorStocks returns the mappings of
// creators and their corresponding stocks
func (s *Server) ListCreatorStocks(ctx context.Context) http.HandlerFunc {
	req := new(listCreatorStockRequest)
	return func(rw http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&req)
		arg := db.ListCreatorStocksParams{
			Limit:  req.Limit,
			Offset: req.Offset,
		}
		creatorStocks, err := s.store.ListCreatorStocks(ctx, arg)
		if err != nil {
			http.Error(rw, "error returning list of creators and their stocks", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(rw).Encode(creatorStocks)
	}
}

// CRUD operations for Creator Porftolio

// GetPortfolioByCreatorID returns a creator portfolio
// provided "creator" id
func (s *Server) GetPortfolioByCreatorID(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)

		creatorPortfolio, err := s.store.GetPortfolioByCreatorID(ctx, id)
		if err != nil {
			http.Error(rw, "error finding the creator portfolio", http.StatusBadRequest)
			return
		}
		json.NewEncoder(rw).Encode(creatorPortfolio)
	}
}

type updateCStockQuantityRequest struct {
	CreatorID int64  `json:"creator_id"`
	StockID   int64  `json:"stock_id"`
	Quantity  string `json:"quantity"`
}

// UpdateCStockQuantity updates the quantity of a
// particular stock in the creator portfolio
func (s *Server) UpdateCStockQuantity(ctx context.Context) http.HandlerFunc {
	req := new(updateCStockQuantityRequest)
	return func(rw http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&req)
		arg := db.UpdateCreatorStockQuantityParams{
			CreatorID: req.CreatorID,
			StockID:   req.StockID,
			Quantity:  req.Quantity,
		}
		if err := s.store.UpdateCreatorStockQuantity(ctx, arg); err != nil {
			http.Error(rw, "error updating the stock quantity", http.StatusInternalServerError)
			return
		}
	}
}

// DeleteStockFromCreatorPortfolio deletes the stocks
// when a user sells them all
func (s *Server) DeleteStockFromCreatorPortfolio(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)
		if err := s.store.DeleteStockFromCreatorPortfolio(ctx, id); err != nil {
			http.Error(rw, "error deleting the stock", http.StatusInternalServerError)
			return
		}
	}
}
