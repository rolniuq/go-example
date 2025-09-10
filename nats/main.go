package main

import "nats/cmd"

func main() {
	if err := cmd.NewApp().Run(); err != nil {
		panic(err)
	}
}
