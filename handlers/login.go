package handlers

import (
	"strings"
	"time"

	"github.com/Attendify-id/auth-services/database"
	"github.com/Attendify-id/auth-services/interfaces"
	"github.com/Attendify-id/auth-services/models"
	"github.com/Attendify-id/auth-services/utils"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	//? Parsing body request ke dalam objek LoginRequest
	json := new(interfaces.LoginRequest)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(interfaces.ResponseJSON{
			Status:  false,
			Message: "Error encoding JSON",
		})
	}

	//? Validasi kredensial
	errResponse := validateCredentials(json)
	if errResponse != nil {
		return c.JSON(*errResponse)
	}

	//? Mencari user dari database berdasarkan username
	foundUser, err := findUserByUsername(json.Username)
	if err != nil {
		return c.JSON(interfaces.ResponseJSON{
			Status:  false,
			Message: "Username not found",
		})
	}

	//? Memverifikasi password
	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(json.Password)); err != nil {
		return c.JSON(interfaces.ResponseJSON{
			Status:  false,
			Message: "Invalid password",
		})
	}

	//? Membuat JWT token
	token, err := utils.CreateJWT(foundUser.Id, foundUser.Username, foundUser.Password)
	if err != nil {
		return c.JSON(interfaces.ResponseJSON{
			Status:  false,
			Message: err.Error(),
		})
	}

	//? Membuat session token di database
	err = createSessionToken(token, foundUser.Id)
	if err != nil {
		return c.JSON(interfaces.ResponseJSON{
			Status:  false,
			Message: "Error creating session token",
		})
	}

	//? Update data user terkait login
	err = updateUserLoginData(c, foundUser)
	if err != nil {
		return c.JSON(interfaces.ResponseJSON{
			Status:  false,
			Message: "Error updating user login data",
		})
	}

	//? Mengembalikan respons JSON dengan data login yang berhasil
	loginData := mapLoginData(foundUser, token)
	return c.JSON(interfaces.ResponseJSON{
		Status:  true,
		Message: "Login successful",
		Data:    loginData,
	})
}

func validateCredentials(json *interfaces.LoginRequest) *interfaces.ResponseJSON {
	credentials := validation.ValidateStruct(json,
		validation.Field(&json.Username, validation.Required, validation.Length(8, 16)),
		validation.Field(&json.Password, validation.Required, validation.Length(8, 16)),
	)

	if credentials != nil {
		return &interfaces.ResponseJSON{
			Status:  false,
			Message: "Error Credentials",
			Error:   credentials,
		}
	}
	return nil
}

func findUserByUsername(username string) (*models.User, error) {
	db := database.DB
	var foundUser models.User
	query := models.User{Username: username}

	if err := db.Preload("UserLevel").Where("username = ?", query.Username).First(&foundUser).Error; err != nil {
		return nil, err
	}
	return &foundUser, nil
}

func createSessionToken(token string, userID int) error {
	db := database.DB
	expires := time.Now().Add(24 * time.Hour)
	dataToken := strings.Split(token, ".")[1]
	session := models.SessionToken{Token: dataToken, ExpiresAt: expires, UserID: userID}
	return db.Create(&session).Error
}

func updateUserLoginData(c *fiber.Ctx, user *models.User) error {
	db := database.DB
	ipAddress := c.IP()
	browser := c.Get("User-Agent")
	now := time.Now()

	user.IpAddress = &ipAddress
	user.Browser = &browser
	user.LastLogin = &now

	return db.Save(user).Error
}

func mapLoginData(user *models.User, token string) fiber.Map {
	return fiber.Map{
		"username":   user.Username,
		"fullname":   user.Fullname,
		"user_level": user.UserLevel.LevelName,
		"token":      token,
	}
}
