package main

import "log"

// use recover to recover from panic
// recover will stop the panic and let it continue with execution

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("recovered", r)
		}
	}()

	myPanic := func() {
		panic("my panic")
	}
	myPanic()

	log.Println("after panic")
}
