package main

import (
	"fmt"
	"time"
)

// select lets you wait on multiple channels operations
// combining goroutines and channels with select is a powerful on go

func main() {
	m := make(chan string)
	n := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		m <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		n <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-m:
			fmt.Println("received", msg1)
		case msg2 := <-n:
			fmt.Println("received", msg2)
		}
	}
}
