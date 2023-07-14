package auth

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	svc authService
}

func newHandler(svc authService) authHandler {
	return authHandler{
		svc: svc,
	}
}

func (a authHandler) signUp(c *fiber.Ctx) error {
	var req register

	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	item, err := a.svc.register(c.UserContext(), req)
	if err != nil {
		log.Println(err)
		return fiber.ErrInternalServerError
	}
	return c.JSON(item)
}
