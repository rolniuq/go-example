package screens

import "fmt"

type TerminalScreen interface {
	Show()
}

type terminal struct{}

func NewTerminalScreen() TerminalScreen {
	return &terminal{}
}

func (t *terminal) Show() {
	fmt.Println(`
		Base64 Apps \n
		Please select the action you want to perform: \n
		1. Encode to string \n
		2. Decode to string
	`)
}
