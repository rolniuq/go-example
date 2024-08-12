package main

import (
	"fmt"
	"time"
)

type Increment struct{}
type GetCount struct {
	Response chan int
}

func counter(incrementChan <-chan Increment, getCountChan <-chan GetCount) {
	count := 0
	for {
		select {
		case <-incrementChan:
			count++
		case get := <-getCountChan:
			get.Response <- count
		}
	}
}

func main() {
	incrementChan := make(chan Increment)
	getCountChan := make(chan GetCount)

	go counter(incrementChan, getCountChan)

	for i := 0; i < 5; i++ {
		incrementChan <- Increment{}
	}

	responseChan := make(chan int)
	getCountChan <- GetCount{Response: responseChan}
	count := <-responseChan

	fmt.Println("Count:", count)

	time.Sleep(1 * time.Second)
}
