package main

import "fmt"

// Map is built in by GO, to associate - ket hop data types
// Sometimes called Hash Or Dict

// map[type]type

func main() {
	m := make(map[string]string)
	m["q"] = "no upgrade"
	fmt.Println(m)
}
