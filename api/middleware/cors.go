package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetCors(app *fiber.App) {
	config := cors.Config{
		AllowOrigins:     "http://127.0.0.1:5500",
		AllowHeaders:     "X-Api-Key,X-User-Id",
		AllowCredentials: true,
	}
	app.Use(cors.New(config))
}
