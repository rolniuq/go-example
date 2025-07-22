package main

import (
	"todo-cli/cmd"
	"todo-cli/configs"
	"todo-cli/ui"
)

func main() {
	config, err := configs.NewConfig().Load()
	if err != nil {
		panic(err)
	}

	app := cmd.NewApp(config)
	if err := app.Start(); err != nil {
		panic(err)
	}

	ui := ui.NewApp(config)
	if err := ui.Start(); err != nil {
		panic(err)
	}
}
