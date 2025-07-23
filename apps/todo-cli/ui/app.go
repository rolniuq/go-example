package ui

import (
	"fmt"
	"todo-cli/configs"
	"todo-cli/pkgs/storage"
	"todo-cli/ui/windows"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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

	res := []fyne.CanvasObject{}
	data := a.storage.GetData()
	for i, item := range data {
		res = append(res, widget.NewLabel(fmt.Sprintf("%d", i)), widget.NewLabel(item.Title), widget.NewLabel(item.DoneString()))
	}

	wd.SetContent(container.New(layout.NewFormLayout(), res...))

	wd.ShowAndRun()

	return nil
}
