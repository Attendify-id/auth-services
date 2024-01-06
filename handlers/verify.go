package handlers

import (
	"strings"
	"time"

	"github.com/Attendify-id/auth-services/database"
	"github.com/Attendify-id/auth-services/models"
	"github.com/Attendify-id/auth-services/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Verify(c *fiber.Ctx) (models.User, error) {
	//? Mendapatkan token dari header Authorization
	authorization := c.Get("Authorization")
	tokens := strings.Split(authorization, " ")[1]

	//? Verifikasi token menggunakan kunci yang sesuai
	token, err := jwt.ParseWithClaims(tokens, &utils.JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return utils.JWT_KEY(), nil
	})
	if err != nil || !token.Valid {
		return models.User{}, fiber.ErrBadRequest
	}

	//? Mendapatkan bagian terenkripsi dari token
	splitter := strings.Split(tokens, ".")[1]
	query := models.SessionToken{Token: splitter}

	//? Mencari sesi token di database
	db := database.DB
	var found models.SessionToken
	if err := db.Where("token = ?", query).First(&found).Error; err != nil {
		return models.User{}, fiber.ErrBadRequest
	}

	//? Memastikan sesi token belum expired di database
	now := time.Now()
	if found.ExpiresAt.Before(now) {
		return models.User{}, fiber.ErrBadRequest
	}

	//? Mengambil data pengguna berdasarkan ID dari sesi token
	var user models.User
	if err := db.Preload("UserLevel").First(&user, found.UserID).Error; err != nil {
		return models.User{}, fiber.ErrBadRequest
	}

	return user, nil
}
