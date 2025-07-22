package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("none")
	name := f.Name()
	if err != nil {
		return
	}

	fmt.Println(name)
}
