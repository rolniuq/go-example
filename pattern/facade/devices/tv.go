package devices

import "fmt"

type TV struct{}

func (t *TV) On() {
	fmt.Println("Turning on the TV...")
}

func (t *TV) SetInput(input string) {
	fmt.Printf("Setting TV input to %s...\n", input)
}
