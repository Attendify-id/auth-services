package handlers

import (
	"github.com/Attendify-id/auth-services/interfaces"
	"github.com/Attendify-id/auth-services/models"
	"github.com/gofiber/fiber/v2"
)

func GetUserInfo(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	return c.JSON(interfaces.ResponseJSON{
		Status:  true,
		Message: "success get user",
		Data:    user,
	})
}
