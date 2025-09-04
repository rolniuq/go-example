package main

import "fmt"

type App struct {
	name string
}

var instance *App

func createApp() *App {
	if instance == nil {
		fmt.Println("App is not exist. Create new")
		instance = &App{
			name: "application",
		}
	}

	return instance
}

func main() {
	app := createApp()
	fmt.Println(app.name)

	app2 := createApp()
	fmt.Println(app2.name)
}
