package controllers

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/firminoneto11/sse-server/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

// This function can be used to create new Controllers
func NewController(connectedClients *shared.ConnectedClients) Controller {
	return Controller{connectedClients: connectedClients}
}

type Controller struct {
	connectedClients *shared.ConnectedClients
}

func (ctr *Controller) SSEHandler(context *fiber.Ctx) error {
	apiKey, userIdStr := context.Query("key"), context.Query("id")

	if apiKey == "" || userIdStr == "" {
		return context.SendString("Invalid headers set.")
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return context.SendString("Invalid headers set.")
	}

	// Upgrading the HTTP connection to be in the SSE format
	context.Set("Content-Type", "text/event-stream")
	context.Set("Cache-Control", "no-cache")
	context.Set("Connection", "keep-alive")
	context.Set("Transfer-Encoding", "chunked")

	// Connecting the client
	ctr.connectedClients.ConnectClient(userId, apiKey)

	// Channel that will receive the data
	clientChannel := ctr.connectedClients.GetClientChannel(userId)
	if clientChannel == nil {
		return context.SendString("Client has disconnected.")
	}

	streamWriter := fasthttp.StreamWriter(
		func(ioWriter *bufio.Writer) {
			fmt.Println("New SSE Connection stablish!")
			for event := range clientChannel {
				// Send the message to the client as a SSE
				message := fmt.Sprintf("event: %s\ndata: "+event.Data+"\n\n", event.Name)
				fmt.Fprint(ioWriter, message)

				err := ioWriter.Flush()
				if err != nil {
					// Refreshing page in web browser will establish a new SSE connection, but only (the last) one is alive, so
					// dead connections must be closed here.
					fmt.Printf("Error while flushing: %v. Closing http connection.\n", err)
					break
				}
			}
			ctr.connectedClients.DisconnectClient(userId)
		},
	)

	// Starts streaming inside this goroutine
	context.Context().SetBodyStreamWriter(streamWriter)

	return nil
}

func (contr *Controller) NewEventHandler(context *fiber.Ctx) error {
	event := shared.Event{Name: "backendTaskReady", Data: "Hey bro!"}
	go contr.connectedClients.SendEvent(10, event)
	return context.SendString("Ok")
}
