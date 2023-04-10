package main

import (
	apiConf "github.com/firminoneto11/sse-server/api/conf"
	"github.com/firminoneto11/sse-server/shared"
)

func main() {
	var connectedClients shared.ConnectedClients
	app := apiConf.GetApp(&connectedClients)
	app.Listen(apiConf.GetPort())
}
