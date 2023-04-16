package routers

import (
	"github.com/firminoneto11/sse-server/api/apps/core/controllers"
	"github.com/firminoneto11/sse-server/shared"
	"github.com/gofiber/fiber/v2"
)

func AddRouter(app *fiber.App, clients *shared.ConnectedClients) {
	router := app.Group("/api")

	controller := controllers.NewController(clients)

	router.Get("/sse/", controller.SSEHandler)
	router.Post("/new-event/", controller.NewEventHandler)
}
