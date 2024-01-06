package utils

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func JWT_KEY() []byte {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	return []byte(os.Getenv("JWT_KEY"))
}

type JWTClaims struct {
	Username string
	Fullname string
	jwt.RegisteredClaims
}

var expires = time.Now().Add(24 * time.Hour)

func CreateJWT(id int, username string, fullname string) (string, error) {

	claims := JWTClaims{
		username,
		fullname,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expires),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "Attendify",
			Subject:   "Easier absences, more regular attendance with Attendify",
			ID:        strconv.Itoa(id),
		},
	}

	createToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := createToken.SignedString(JWT_KEY())

	return token, err
}
