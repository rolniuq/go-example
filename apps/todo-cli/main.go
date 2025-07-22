package main

import (
	"todo-cli/cmd"
	"todo-cli/configs"
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
}
