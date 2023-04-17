package controllers

import (
	"bufio"
	"fmt"
	"strconv"
	"time"

	"github.com/firminoneto11/sse-server/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func NewController(clients *shared.ConnectedClients) Controller {
	return Controller{clients: clients}
}

type Controller struct {
	clients *shared.ConnectedClients
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

	// Channel that will receive the data
	clientChannel := make(chan shared.Event)

	// Connecting this client
	ctr.clients.ConnectClient(userId, &clientChannel)

	disconnectChan := func() {
		ctr.clients.DisconnectClient(userId, &clientChannel)
		_, ok := <-clientChannel
		if ok {
			close(clientChannel)
		}
	}

	streamWriter := fasthttp.StreamWriter(
		func(ioWriter *bufio.Writer) {

			go func() {
				for {
					fmt.Fprintf(ioWriter, "event: heartbeat\n\n")
					err := ioWriter.Flush()
					if err != nil {
						disconnectChan()
						break
					}
					time.Sleep(time.Second)
				}
			}()

			for event := range clientChannel {
				message := fmt.Sprintf("event: %s\ndata: "+event.Data+"\n\n", event.Name)
				fmt.Fprint(ioWriter, message)
				err := ioWriter.Flush()
				if err != nil {
					disconnectChan()
					break
				}
			}
		},
	)

	// Starts streaming inside this goroutine
	context.Context().SetBodyStreamWriter(streamWriter)

	return nil
}

func (ctr *Controller) NewEventHandler(context *fiber.Ctx) error {
	event := shared.Event{Name: "backendTaskReady", Data: "Hey bro!"}

	go func() {
		for {
			ctr.clients.BroadCastEvent(10, event)
			time.Sleep(time.Second)
		}
	}()

	return context.SendString("Ok")
}
