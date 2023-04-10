package conf

import (
	apiMiddleware "github.com/firminoneto11/sse-server/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func GetApp() *fiber.App {
	app := fiber.New()
	apiMiddleware.SetCors(app)
	setRouters(app)
	return app
}

func GetPort() string {
	// TODO: Make this as an env variable set in the 'settings' file
	return ":8000"
}
