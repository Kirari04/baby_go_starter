package main

import (
	"baby_starter/app"
	"baby_starter/database"
	"baby_starter/server"
)

func main() {
	app.Init()
	database.Init()

	server.Start()
}
