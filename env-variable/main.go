package main

import (
	"fmt"
	"os"
	"strings"
)

// a universal mechanism for conveying configuration information to Unix programs

func main() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO", os.Getenv("FOO"))
	fmt.Println("BAR", os.Getenv("BAR"))

	fmt.Println()
	for _, v := range os.Environ() {
		splitN := strings.SplitN(v, "=", 2)
		// -> FOO=1
		// -> [FOO, 1]
		res := splitN[0] // FOO
		fmt.Println(res)
	}
}

// normal run command: go run main.go
// want to run with env from terminal: BAR=2 go run main.go
