package post_delivery

import (
	"texts/internal/middleware"
	"texts/pkg/constants"

	"github.com/gofiber/fiber/v2"
)

func MapPostRoutes(group fiber.Router, h Handlers, mw *middleware.MDWManager) {
	group.Post("/create_post", h.CreatePost())
	group.Post("/get_posts", h.GetPosts())
	group.Post("/approve_post", mw.AuthedMiddleware(&constants.RoleEditor), h.ApprovePost())
	group.Post("/reject_post", mw.AuthedMiddleware(&constants.RoleEditor), h.RejectPost())
}
