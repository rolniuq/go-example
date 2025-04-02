package main

import "fmt"

// create channel unbuffer channel and buffer channel
// launch few goroutines that send value to channel
// main go routine receive value from channel and print
// demonstrate how them works
// use select for multiple channel

type Lesson1 struct{}

func (l *Lesson1) Execute() {
	unCh := make(chan int)
	bufCh := make(chan int, 3)
	done := make(chan bool)

	defer func() {
		close(unCh)
		close(bufCh)
		close(done)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			unCh <- i
			bufCh <- i
		}

		done <- true
	}()

	for {
		select {
		case v := <-unCh:
			fmt.Println("unbuffered channel", v)
		case v := <-bufCh:
			fmt.Println("buffered channel", v)
		case <-done:
			return
		}
	}
}

func (l *Lesson1) Fixed() {
	unCh := make(chan int)
	bufCh := make(chan int, 3)
	done := make(chan struct{})

	go func() {
		for i := 0; i < 10; i++ {
			unCh <- i
			bufCh <- i
		}

		close(done)
		close(unCh)
		close(bufCh)
	}()

	for {
		select {
		case v, ok := <-unCh:
			if !ok {
				unCh = nil // Avoid further reads from closed channel
			} else {
				fmt.Println("unbuffered channel:", v)
			}
		case v, ok := <-bufCh:
			if !ok {
				bufCh = nil // Avoid further reads from closed channel
			} else {
				fmt.Println("buffered channel:", v)
			}
		case <-done:
			return
		}
	}
}
