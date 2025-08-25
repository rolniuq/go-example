package screens

import "base64parser/ui/screens"

type TerminalScreen interface {
	Show()
}

type terminal struct {
	ui screens.TerminalScreen
}

func NewTerminalScreen() TerminalScreen {
	return &terminal{}
}

func (t *terminal) Show() {
	t.ui.Show()
}
