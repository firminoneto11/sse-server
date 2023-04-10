package main

import apiConf "github.com/firminoneto11/sse-server/api/conf"

func main() {
	app := apiConf.GetApp()
	app.Listen(apiConf.GetPort())
}
