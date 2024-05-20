package main

import (
	"github.com/GinoCodeSpace/bridge/config"
	"github.com/GinoCodeSpace/bridge/router"
)

func main() {
	config.Init()

	router.Initialize()
}
