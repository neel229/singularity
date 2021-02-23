package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	db "github.com/neel229/singularity/pkg/db/sqlc"
)

// Server contains an instance to fiber.App
// and a store.
type Server struct {
	app   *fiber.App
	store *db.Store
}

// NewServer creates a new server
func NewServer(s *db.Store) *Server {
	return &Server{
		app:   fiber.New(),
		store: s,
	}
}

// StartServer starts a new server on
// port 69420
func (s *Server) StartServer(addr string) {
	if err := s.app.Listen(":5000"); err != nil {
		log.Fatalf("error starting the server: %v", err)
	}
}
