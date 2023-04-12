package conf

import (
	coreRouters "github.com/firminoneto11/sse-server/api/apps/core/routers"
	apiMiddleware "github.com/firminoneto11/sse-server/api/middleware"
	"github.com/firminoneto11/sse-server/shared"
	"github.com/gofiber/fiber/v2"
)

func GetApp(connectedClients *shared.ConnectedClients) *fiber.App {
	app := fiber.New()
	apiMiddleware.SetCors(app)

	// NOTE: The amount of routers for each app is likely to increase, so you should refactor this later to a function that sets
	// routers for example.
	coreRouters.AddRouter(app, connectedClients)

	return app
}

func GetPort() string {
	// TODO: Make this as an env variable set in the 'settings' file
	return ":8007"
}
