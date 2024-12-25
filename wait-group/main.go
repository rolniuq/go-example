package main

import (
	"fmt"
	"sync"
	"time"
)

// to wait for multiple goroutines finish we use a wait group

func worker(id int) {
	fmt.Println("worker", id)
	time.Sleep(2 * time.Second)
	fmt.Println("worker", id, "finished")
}

func Doc() {
	// the wait group is used to wait for all the goroutines launched to finish
	// if a wait group is explicitly passed into functions, it should be done by pointer
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		// launch several goroutines and increment the wait group counter for each
		wg.Add(1)

		// wrap the worker call in a closure that makes sure to tell the wait group that this worker is done
		// this way the worker itself does not have to be aware of the concurrency primitive involved in its execution
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
}

// You can guarantee the goroutine count is fixed and correct.
// The performance of your program is critical.
func AddOnceTime() {
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
}

// The goroutine count is dynamic or unclear at compile time.
// You want safer, more maintainable code that avoids potential bugs related to mismatched Add calls.
func AddMultipleTimes() {
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
}

func main() {
	Doc()
	AddOnceTime()
	AddMultipleTimes()
}
