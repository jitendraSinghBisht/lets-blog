package main

import (

	"lets-blog/config"
)

func main() {
	config.Init()
	serverConfig := config.GetInstance()

	app := NewApp(serverConfig)
	app.Start()
}
