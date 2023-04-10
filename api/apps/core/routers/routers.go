package routers

import (
	"github.com/firminoneto11/sse-server/api/apps/core/controllers"
	"github.com/gofiber/fiber/v2"
)

func AddRouter(app *fiber.App) {
	router := app.Group("/sse")
	router.Get("/", controllers.Hello)
}
