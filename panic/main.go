package main

import "log"

// mostly we want to use panic when we want fail fast
// when use panic it will drop all the stack and crash the program

func main() {
	panic("something went wrong")
	log.Println("after panic")
}
