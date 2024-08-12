package main

import (
	"fmt"
	"time"
)

// Times are for when we want to do something in the future
// Tickers are for when we want to do something repeated at regular intervals

func main() {
	// tickers use a similar mechanism to timers
	// a channel that is sent values
	//we will use the select builtin on the channel to await the values as they arrive every 500 ms
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// ticker can be stopped like timers.
	// once a ticker is stopped it wont receive any more value on its channel
	// we will stop our after 1600 ms
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
