package conf

import (
	coreRouters "github.com/firminoneto11/sse-server/api/apps/core/routers"
	"github.com/gofiber/fiber/v2"
)

func setRouters(app *fiber.App) {
	coreRouters.AddRouter(app)
}
