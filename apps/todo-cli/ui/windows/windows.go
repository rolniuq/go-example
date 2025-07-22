package windows

import (
	"todo-cli/configs"

	"fyne.io/fyne/v2"
)

type (
	WindowContent fyne.CanvasObject

	Windows struct {
		fyne.Window
	}
)

func NewWindows(ui fyne.App, config *configs.Config) *Windows {
	w := ui.NewWindow(config.AppName)

	return &Windows{w}
}

func (w *Windows) SetContent(content WindowContent) {
	w.Window.SetContent(content)
}

func (w *Windows) ShowAndRun() {
	if w == nil {
		return
	}

	w.Window.ShowAndRun()
}
