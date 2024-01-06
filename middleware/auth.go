package middleware

import (
	"github.com/Attendify-id/auth-services/handlers"
	"github.com/Attendify-id/auth-services/interfaces"
	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	user, err := handlers.Verify(c)
	if err != nil {
		return c.JSON(interfaces.ResponseJSON{
			Status:  false,
			Message: "unauthorized",
		})
	}
	c.Locals("user", user)
	return c.Next()
}
