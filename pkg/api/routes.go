package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
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
	s.r.Route("/currency", func(r chi.Router) {
		r.Get("/", s.ListCurrencies(s.ctx))
		r.Post("/", s.CreateCurrency(s.ctx))

		// Subroutes for currency
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", s.GetCurrency(s.ctx))
			r.Put("/", s.UpdateCurrency(s.ctx))
			r.Delete("/", s.DeleteCurrency(s.ctx))
		})
	})
}

func homepage(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("This is the homepage")
}
