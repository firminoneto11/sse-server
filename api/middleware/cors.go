package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetCors(app *fiber.App) {
	config := cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}
	app.Use(cors.New(config))
}
