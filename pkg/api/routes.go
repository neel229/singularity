package api

import (
	"context"
	"encoding/json"
	"net/http"
)

// SetRoutes sets the routes with
// corresponding handler functions
func (s *Server) SetRoutes() {
	s.r.Get("/home", homepage)
	s.r.Post("/create/creator", s.CreateCreator(context.Background()))
}

func homepage(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("This is the homepage")
}
