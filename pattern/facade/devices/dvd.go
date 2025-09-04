package devices

import "fmt"

type DVDPlayer struct{}

func (d *DVDPlayer) On() {
	fmt.Println("Turning on the DVD Player...")
}

func (d *DVDPlayer) Play(movie string) {
	fmt.Printf("Playing movie '%s'...\n", movie)
}
