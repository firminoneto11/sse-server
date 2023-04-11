package main

import (
	apiConf "github.com/firminoneto11/sse-server/api/conf"
	"github.com/firminoneto11/sse-server/shared"
)

func main() {
	connectedClients := shared.NewConnectedClients()
	app := apiConf.GetApp(&connectedClients)
	app.Listen(apiConf.GetPort())
}
