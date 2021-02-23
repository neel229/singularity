package api

import (
	"encoding/json"
	"net/http"
)

// SetRoutes sets the routes with
// corresponding handler functions
func (s *Server) SetRoutes() {
	s.r.Get("/home", homepage)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("This is the homepage")
}

// // Logic for creating a creator account
// // and corresponding stock
// type createCreatorRequest struct {
// 	FirstName           string `json:"first_name" validate:"required"`
// 	LastName            string `json:"last_name" validate:"required"`
// 	UserName            string `json:"user_name" validate:"required"`
// 	Email               string `json:"email" validate:"required, email"`
// 	Password            string `json:"password" validate:"required"`
// 	PreferredCurrencyID int32  `json:"preferred_currency_id" validate:"required"`
// 	Ticker              string `json:"ticker" validate:"required"`
// 	Details             string `json:"details" validate:"required"`
// }

// func createCreator(c *fiber.Ctx) error {
// 	req := new(createCreatorRequest)
// 	if err := c.BodyParser(req); err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
// 	}

// 	arg := db.StockCreationTxParams{
// 		FirstName:           req.FirstName,
// 		LastName:            req.LastName,
// 		UserName:            req.UserName,
// 		Email:               req.Email,
// 		Password:            req.Password,
// 		PreferredCurrencyID: req.PreferredCurrencyID,
// 		Ticker:              req.Ticker,
// 		Details:             req.Details,
// 	}

// 	s := Server{
// 		store: NewStore(),
// 	}
// 	return nil
// }
