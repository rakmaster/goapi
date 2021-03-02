package main

import (
	"log"

	"github.com/rakmaster/goapi/app"
	"github.com/rakmaster/goapi/config"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	config := config.NewConfig()
	app.ConfigAndRunApp(config)
}
