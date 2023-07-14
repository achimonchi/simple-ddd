package auth

import (
	"github.com/gofiber/fiber/v2"
)

func Run(router fiber.Router) {
	memmoryRepo := newMemmory()
	svc := newService(memmoryRepo)
	handler := newHandler(svc)
	auth := router.Group("auth")
	{
		auth.Post("/signup", handler.signUp)
		// auth.Post("/signin")
	}
}
