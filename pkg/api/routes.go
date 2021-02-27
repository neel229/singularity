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
	s.vofferRoutes()
	s.vtradeRoutes()
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
			r.Route("/portfolio", func(r chi.Router) {
				r.Get("/", s.GetFanPortfolio(s.ctx))
				r.Put("/", s.UpdateFanStockQuantity(s.ctx))
				r.Delete("/", s.DeleteStockFromFanPortfolio(s.ctx))
			})
		})
	})
}

// Setup all the routes surrounding a creator
func (s *Server) creatorRoutes() {
	s.r.Route("/creator", func(r chi.Router) {
		r.Get("/", s.ListCreators(s.ctx))
		r.Post("/", s.CreateCreatorAndStock(s.ctx))

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", s.GetCreator(s.ctx))
			r.Get("/vtokens", s.GetVirginTokensLeft(s.ctx))
			r.Get("/voffer", s.GetVirginOfferByCreator(s.ctx))
			r.Put("/email", s.UpdateCreatorEmail(s.ctx))
			r.Put("/password", s.UpdateCreatorPassword(s.ctx))
			r.Put("/pcurrency", s.UpdateCreatorPreferredCurrency(s.ctx))
			r.Put("/vtokens", s.UpdateCreatorVTLeft(s.ctx))
			r.Route("/portfolio", func(r chi.Router) {
				r.Get("/", s.GetPortfolioByCreatorID(s.ctx))
				r.Put("/", s.UpdateCStockQuantity(s.ctx))
				r.Delete("/", s.DeleteStockFromCreatorPortfolio(s.ctx))
			})
		})
		r.Route("/stock", func(r chi.Router) {
			r.Get("/", s.ListCreatorStocks(s.ctx))
			r.Get("/{id}", s.GetCreatorStock(s.ctx))
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

func (s *Server) vofferRoutes() {
	s.r.Route("/voffers", func(r chi.Router) {
		r.Get("/", s.ListVirginOffers(s.ctx))
		r.Post("/", s.CreateVirginOffer(s.ctx))

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", s.GetVirginOffer(s.ctx))
		})
	})
}

func (s *Server) vtradeRoutes() {
	s.r.Route("/vtrade", func(r chi.Router) {
		r.Post("/", s.VTradeTx(s.ctx))
	})
}
