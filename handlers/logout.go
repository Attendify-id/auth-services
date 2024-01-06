package handlers

import (
	"strings"
	"time"

	"github.com/Attendify-id/auth-services/database"
	"github.com/Attendify-id/auth-services/interfaces"
	"github.com/Attendify-id/auth-services/models"
	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {
	db := database.DB
	authorization := c.Get("Authorization")
	sessionKey := strings.Split(authorization, " ")[1]

	query := models.SessionToken{Token: sessionKey}
	found := models.SessionToken{}

	if err := db.Preload("UserLevel").Where("token = ?", query.Token).First(&found).Error; err != nil {
		return c.JSON(interfaces.ResponseJSON{
			Status:  false,
			Message: "token not found",
		})
	}

	// now := time.Now()
	found.ExpiresAt = time.Now()
	db.Save(&found)

	return c.JSON(interfaces.ResponseJSON{
		Status:  true,
		Message: "logout success",
	})
}
