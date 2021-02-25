package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// SetRoutes sets the routes with
// corresponding handler functions
func (s *Server) SetRoutes() {
	s.r.Use(middleware.Logger)
	s.currencyRoutes()
	s.currencyRateRoutes()
	s.fanRoutes()
	s.creatorRoutes()
	s.stockRoutes()
	s.creatorStockRoutes()
	s.creatorPortfolioRoutes()
}

// Sets up routes for currency
func (s *Server) currencyRoutes() {
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

// Sets up routes for currency rates
func (s *Server) currencyRateRoutes() {
	s.r.Route("/currency/rate", func(r chi.Router) {
		r.Post("/", s.CreateCurrencyRate(s.ctx))

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", s.GetCurrencyRate(s.ctx))
			r.Put("/", s.UpdateCurrencyRate(s.ctx))
		})
	})
}

// Sets up routes for fan
func (s *Server) fanRoutes() {
	s.r.Route("/fan", func(r chi.Router) {
		r.Post("/", s.CreateFan(s.ctx))

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", s.GetFan(s.ctx))
			r.Put("/email", s.UpdateFanEmail(s.ctx))
			r.Put("/password", s.UpdateFanPassword(s.ctx))
			r.Put("/pcurrency", s.UpdateFanPreferredCurrency(s.ctx))
		})
	})
}

// Sets up routes for creator and corresponding stock
func (s *Server) creatorRoutes() {
	s.r.Route("/creator", func(r chi.Router) {
		r.Get("/", s.ListCreators(s.ctx))
		r.Post("/", s.CreateCreatorAndStock(s.ctx))

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", s.GetCreator(s.ctx))
			r.Get("/vtokens", s.GetVirginTokensLeft(s.ctx))
			r.Get("/portfolio", s.GetPortfolioByCreatorID(s.ctx))
			r.Put("/email", s.UpdateCreatorEmail(s.ctx))
			r.Put("/password", s.UpdateCreatorPassword(s.ctx))
			r.Put("/pcurrency", s.UpdateCreatorPreferredCurrency(s.ctx))
			r.Put("/vtokens", s.UpdateCreatorVTLeft(s.ctx))
		})
	})
}

func (s *Server) stockRoutes() {
	s.r.Route("/stock", func(r chi.Router) {
		r.Get("/", s.ListStocks(s.ctx))

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", s.GetStock(s.ctx))
			r.Put("/", s.UpdateStockDetails(s.ctx))
		})
	})
}

func (s *Server) creatorStockRoutes() {
	s.r.Route("/creator/stock", func(r chi.Router) {
		r.Get("/", s.ListCreatorStocks(s.ctx))
		r.Get("/{id}", s.GetCreatorStock(s.ctx))
	})
}

func (s *Server) creatorPortfolioRoutes() {
	s.r.Put("/portfolio", s.UpdateCStockQuantity(s.ctx))
	s.r.Delete("/portfolio/{id}", s.DeleteStockFromCreatorPortfolio(s.ctx))
}
