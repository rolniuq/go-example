package ui

import (
	"fyne.io/fyne/v2/app"
)

type UI struct{}

func NewUI() *UI {
	return &UI{}
}

func (u *UI) Start() {
	a := app.New()
	w := a.NewWindow("Base64")

	w.ShowAndRun()
}
