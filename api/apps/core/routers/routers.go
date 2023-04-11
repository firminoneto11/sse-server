package routers

import (
	"github.com/firminoneto11/sse-server/api/apps/core/controllers"
	"github.com/firminoneto11/sse-server/shared"
	"github.com/gofiber/fiber/v2"
)

func AddRouter(app *fiber.App, connectedClients *shared.ConnectedClients) {
	router := app.Group("/api")

	controller := controllers.NewController(connectedClients)

	router.Get("/sse/", controller.SSEHandler)
	router.Post("/new-event/", controller.NewEventHandler)
}
