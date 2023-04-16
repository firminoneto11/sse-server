package main

import (
	apiConf "github.com/firminoneto11/sse-server/api/conf"
	"github.com/firminoneto11/sse-server/shared"
)

func main() {
	clients := shared.NewConnectedClients()
	app := apiConf.GetApp(&clients)
	app.Listen(apiConf.GetPort())
}
