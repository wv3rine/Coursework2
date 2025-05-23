package commentary_delivery

import (
	"texts/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func MapCommentaryRoutes(group fiber.Router, h Handlers, mw *middleware.MDWManager) {
	group.Post("/create_commentary", h.CreateCommentary())
	group.Post("/get_commentary", h.GetCommentaries())
}
