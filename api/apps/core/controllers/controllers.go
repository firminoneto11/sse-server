package controllers

import (
	"fmt"

	"github.com/firminoneto11/sse-server/shared"
	"github.com/gofiber/fiber/v2"
)

// This function can be used to create new Controllers
func NewController(connectedClients *shared.ConnectedClients) Controller {
	return Controller{connectedClients: connectedClients}
}

type Controller struct {
	connectedClients *shared.ConnectedClients
}

func (c *Controller) SSEHandler(context *fiber.Ctx) error {

	connectedClients := c.connectedClients

	connectedClients.AddClient(1, "hello!")
	fmt.Println(connectedClients.IsConnected(1))

	response := "Hello World!"
	return context.SendString(response)
}

func (c *Controller) NewEventHandler(context *fiber.Ctx) error {
	response := "Hello World!"
	return context.SendString(response)
}
