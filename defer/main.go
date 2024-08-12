package main

import "os"

// use defer to ensure that a function call gets executed at the end of the enclosing function

func main() {
	_, err := os.Create("log.txt")
	if err != nil {
		panic(err)
	}
	defer os.Remove("log.txt")
}
