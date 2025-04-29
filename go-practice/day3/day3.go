package day3

import (
	"fmt"
	"time"
)

type Day3 struct{}

func createRoutine(times, ms int, msg string) {
	for range times {
		if ms > 0 {
			duration := time.Duration(ms)
			time.Sleep(duration * time.Millisecond)
		}

		fmt.Println(msg)
	}
}

func (d *Day3) First() {
	go createRoutine(5, 100, "Goroutine 1")
	go createRoutine(3, 200, "Goroutine 2")
	go createRoutine(1, 0, "Goroutine 3")

	fmt.Println("All done")
}

func (d *Day3) Second() {
	ch := make(chan string)

	go func() {
		for range 3 {
			time.Sleep(100 * time.Millisecond)
			ch <- "ping"
		}

		defer close(ch)
	}()

	go func() {
		for c := range ch {
			fmt.Printf("Received: %s\n", c)
		}
	}()

	// for prevent channel close
	time.Sleep(2 * time.Second)
}

func (d *Day3) Third() {
	fastChan := make(chan string)
	defer close(fastChan)
	slowChan := make(chan string)
	defer close(slowChan)

	go func() {
		time.Sleep(100 * time.Millisecond)
		fastChan <- "fast"
	}()

	go func() {
		time.Sleep(time.Second)
		slowChan <- "slow"
	}()

	select {
	case msg := <-fastChan:
		fmt.Println(msg)
	case msg := <-slowChan:
		fmt.Println(msg)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("timeout")
		return
	}
}

func (d *Day3) Exec() {
	d.First()
	d.Second()
	d.Third()
}
