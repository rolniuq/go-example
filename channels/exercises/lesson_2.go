package main

import (
	"fmt"
	"time"
)

/*
  Rules:
  1. No waitgroup
  2. No global variables
  3. Must use select
  4. Must close channel properly

  Requirements:
  1. Create generateNumbers(n int) <-chan int that send n number (0 to n - 1) to a channel
  2. Create worker pool of 3 go routines that receive numbers and square them
  3. Merge results from all workers into one result channel and print them
  4. No waitgroup and select is allowed
*/

func generateNumbers(n int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for i := 0; i < n; i++ {
			ch <- i
		}
	}()

	return ch
}

func work(id int, workers <-chan int, result chan<- int, done chan<- bool) {
	for n := range workers {
		fmt.Printf("worker %d received %d\n", id, n)
		time.Sleep(time.Second)

		result <- n * n
	}

	done <- true
}

type WorkerPool struct {
	number        int
	numberWorkers int

	result chan int
	done   chan bool
}

func NewWorkerPool(n, w int) *WorkerPool {
	return &WorkerPool{
		number:        n,
		numberWorkers: w,

		result: make(chan int, n),
		done:   make(chan bool, w),
	}
}

func (w *WorkerPool) Execute() {
	numbers := generateNumbers(w.number)

	for i := 0; i < w.numberWorkers; i++ {
		go func() {
			work(i, numbers, w.result, w.done)
		}()
	}

	go func() {
		for i := 0; i < w.numberWorkers; i++ {
			<-w.done
		}

		close(w.result)
	}()

	for res := range w.result {
		fmt.Println("Res:", res)
	}
}
