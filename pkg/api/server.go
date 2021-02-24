package api

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	db "github.com/neel229/singularity/pkg/db/sqlc"
)

// Server contains an instance to fiber.App
// and a store.
type Server struct {
	r     *chi.Mux
	store *db.Store
	ctx   context.Context
}

// NewServer creates a new server
func NewServer(s *db.Store) *Server {
	return &Server{
		r:     chi.NewRouter(),
		store: s,
		ctx:   context.Background(),
	}
}

// StartServer starts a server
// on address provided
func (s *Server) StartServer(addr string) {
	if err := http.ListenAndServe(addr, s.r); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
