package user_delivery

import (
	"texts/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func MapUserRoutes(group fiber.Router, h Handlers, mw *middleware.MDWManager) {
	group.Post("/sign_up", mw.NonAuthedMiddleware(), h.SignUp())
	group.Post("/sign_in", mw.NonAuthedMiddleware(), h.SignIn())
	group.Post("/sign_out", mw.AuthedMiddleware(nil), h.SignOut())
	group.Get("/check_session", mw.AuthedMiddleware(nil), h.CheckSession())
}
