package api

import "github.com/gofiber/fiber/v2"

// SetRoutes sets the routes with
// corresponding handler functions
func (s *Server) SetRoutes() {
	s.app.Get("/", homepage)
}

func homepage(c *fiber.Ctx) error {
	c.Send([]byte("This is the homepage"))
	return nil
}
