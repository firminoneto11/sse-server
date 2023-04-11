package controllers

import (
	"fmt"
	"strconv"
	"time"

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

func (contr *Controller) SSEHandler(context *fiber.Ctx) error {
	const eventName string = "backendTaskReady"

	apiKey, userIdStr := context.Query("key"), context.Query("id")

	if apiKey == "" || userIdStr == "" {
		fmt.Println("1")
		return context.SendString("Invalid headers set.")
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		fmt.Println("2")
		return context.SendString("Invalid headers set.")
	}

	// Upgrading the HTTP connection to be in the SSE format
	context.Set("Content-Type", "text/event-stream")
	context.Set("Cache-Control", "no-cache")
	context.Set("Connection", "keep-alive")

	// Connecting the client and deferring its disconnection
	contr.connectedClients.ConnectClient(userId, apiKey)
	defer contr.connectedClients.DisconnectClient(userId)

	// Channel that will receive the data
	clientChannel := contr.connectedClients.GetClientChannel(userId)
	if clientChannel == nil {
		fmt.Println("3")
		return context.SendString("Client has disconnected.")
	}

	// Channel that will be used to signal when the client has disconnected
	connectedChannel := make(chan struct{})

	// Goroutine that listens to the client's events channel and sends them to the browser
	go func() {
		defer close(connectedChannel)
		for eventData := range clientChannel {
			// Send the message to the client as a SSE
			response := fmt.Sprintf("event: %s\ndata: "+eventData+"\n\n", eventName)
			context.SendString(response)
		}
	}()

	// Just for testing...
	go func() {
		for {
			clientChannel <- "Hey mr..."
			time.Sleep(time.Second)
		}
	}()

	select {
	case <-connectedChannel:
		fmt.Println("Client disconnected")
	case <-context.Context().Done():
		fmt.Println("Request cancelled")
	}

	return nil
}

func (contr *Controller) NewEventHandler(context *fiber.Ctx) error {
	response := "Hello World!"
	return context.SendString(response)
}
