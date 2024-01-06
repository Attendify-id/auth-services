package main

import (
	"log"
	"os"

	"github.com/Attendify-id/auth-services/database"
	"github.com/Attendify-id/auth-services/handlers"
	"github.com/Attendify-id/auth-services/middleware"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.ConnectDB()
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Post("/", handlers.Login)
	app.Delete("/", middleware.Auth, handlers.Logout)
	app.Get("/", middleware.Auth, handlers.GetUserInfo)
	panic(app.Listen(":" + os.Getenv("PORT") + ""))
}
