package main

import (
	"fmt"
	"time"
)

// rate limiting is an important mechanism for controlling resource utilization and maintaining quality of service
// go elegantly supports rate limiting with goroutines, channels, and tickers

func main() {
	// suppose we want to limit our handling of incoming request
	// we will serve these requests off a channel of the same name
	requests := make(chan int, 5)
	for i := 1; i < 6; i++ {
		requests <- i
	}
	close(requests)

	// the limiter channel will receive a value every 200 milliseconds
	// this is the regulator in our rate limiting scheme
	limiter := time.Tick(200 * time.Millisecond)

	// by blocking on a receive from the limiter channel before serving each request
	// we limit ourselves to 1 request every 200 milliseconds
	for request := range requests {
		<-limiter
		fmt.Println("request", request, time.Now())
	}

	// we may want to allow short bursts of requests in our rate limiting scheme while preserving the overall
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
