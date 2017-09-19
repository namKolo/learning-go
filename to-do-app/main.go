package main

import (
	"learning-go/to-do-app/config"
)

func main() {
	config := config.GetConfig()

	app := &App{}
	app.Initialize(config)
	app.Run(":3000")
}
