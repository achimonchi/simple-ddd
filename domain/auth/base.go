package auth

import (
	"github.com/gofiber/fiber/v2"
)

func Run(router fiber.Router) {
	svc := newService()
	handler := newHandler(svc)
	auth := router.Group("auth")
	{
		auth.Post("/signup", handler.signUp)
		// auth.Post("/signin")
	}
}
