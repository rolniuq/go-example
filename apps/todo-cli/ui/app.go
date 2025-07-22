package ui

import (
	"todo-cli/configs"
	"todo-cli/pkgs/storage"
	"todo-cli/ui/windows"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	config  *configs.Config
	storage *storage.Storage
}

func NewApp(config *configs.Config) *App {
	storage := storage.NewStorage(config)
	return &App{
		config:  config,
		storage: storage,
	}
}

func (a *App) Start() error {
	ui := app.New()
	wd := windows.NewWindows(ui, a.config)

	hello := widget.NewLabel("Hello Fyne!")
	wd.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	wd.ShowAndRun()

	return nil
}
