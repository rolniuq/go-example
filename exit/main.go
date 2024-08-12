package main

import (
	"fmt"
	"os"
)

// Use os.Exit to immediately exit with a given status.

func main() {
	// defers will not be run when using os.Exit, so this fmt.Println will never be called
	defer fmt.Println("!")

	os.Exit(3)
}
