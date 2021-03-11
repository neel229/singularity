package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	db "github.com/neel229/forum/pkg/db/sqlc"
)

type createVOfferRequest struct {
	CreatorID int64  `json:"creator_id"`
	StockID   int64  `json:"stock_id"`
	Quantity  string `json:"quantity"`
	Price     string `json:"price"`
}

// CreateVirginOffer creates a virgin offer
// between a creator and a fan to initiate a trade
func (s *Server) CreateVirginOffer(ctx context.Context) http.HandlerFunc {
	req := new(createVOfferRequest)
	return func(rw http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&req)
		arg := db.CreateVirginOfferParams{
			CreatorID: req.CreatorID,
			StockID:   req.StockID,
			Quantity:  req.Quantity,
			Price:     req.Price,
		}
		virginOffer, err := s.store.CreateVirginOffer(ctx, arg)
		if err != nil {
			http.Error(rw, "error creating a virign offer", http.StatusBadRequest)
			return
		}
		json.NewEncoder(rw).Encode(virginOffer)
	}
}

// GetVirginOffer fetches a virgin offer from a given id
func (s *Server) GetVirginOffer(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)
		virginOffer, err := s.store.GetVirginOffer(ctx, id)
		if err != nil {
			http.Error(rw, "offer not found, check id", http.StatusBadRequest)
			return
		}
		json.NewEncoder(rw).Encode(virginOffer)
	}
}

// GetVirginOfferByCreator fetches a virgin offer
// from a given creator id
func (s *Server) GetVirginOfferByCreator(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		param, _ := strconv.Atoi(chi.URLParam(r, "id"))
		id := int64(param)
		virginOffer, err := s.store.GetVirginOfferByCreator(ctx, id)
		if err != nil {
			http.Error(rw, "offer not found, check id", http.StatusBadRequest)
			return
		}
		json.NewEncoder(rw).Encode(virginOffer)
	}
}

// ListVirginOffers returns a list of all the virgin offers
// currently present on the platform
func (s *Server) ListVirginOffers(ctx context.Context) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		param, _ := strconv.Atoi(chi.URLParam(r, "limit"))
		param2, _ := strconv.Atoi(chi.URLParam(r, "offset"))
		limit := int32(param)
		offset := int32(param2)
		arg := db.ListVirginOffersParams{
			Limit:  limit,
			Offset: offset,
		}
		offers, err := s.store.ListVirginOffers(ctx, arg)
		if err != nil {
			http.Error(rw, "error returning list of offers", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(rw).Encode(offers)
	}
}
