package ui

import (
	"todo-cli/configs"
	"todo-cli/pkgs/storage"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	storage *storage.Storage
}

func NewApp(config *configs.Config) *App {
	storage := storage.NewStorage(config)
	return &App{
		storage: storage,
	}
}

func (a *App) Start() error {
	ui := app.New()

	w := ui.NewWindow("Todo Cli")
	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()

	return nil
}
