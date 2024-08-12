package main

import (
	"fmt"
	"time"
)

// channel is the pipe which goroutine comunication
// msg <- "hello world"

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(2 * time.Second)
	fmt.Println("done")

	done <- true
}

func channelSynchronization() {
	done := make(chan bool, 1)

	go worker(done)

	<-done
}

func basicChannel() {
	msg := make(chan string)

	go func() {
		msg <- "hello world"
		msg <- "hello again"
	}()

	x := <-msg
	y := <-msg

	fmt.Println(x)
	fmt.Println(y)
}

func main() {
	// basicChannel()
	channelSynchronization()
}
