package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/middleware"
)

// SetRoutes sets the routes with
// corresponding handler functions
func (s *Server) SetRoutes() {
	s.r.Use(middleware.Logger)
	s.r.Get("/home", homepage)
	s.r.Post("/create/creator", s.CreateCreator(s.ctx))
	s.CurrencyRoutes()
}

// CurrencyRoutes sets the routes
// respect to currencies
func (s *Server) CurrencyRoutes() {
	s.r.Get("/currency/{id}", s.GetCurrency(s.ctx))
	s.r.Get("/currencies", s.ListCurrencies(s.ctx))
	s.r.Post("/create/currency", s.CreateCurrency(s.ctx))
	s.r.Put("/update/currency/{id}", s.UpdateCurrency(s.ctx))
	s.r.Get("/currency/{id}", s.GetCurrency(s.ctx))
}

func homepage(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("This is the homepage")
}
