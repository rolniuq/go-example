package main

import (
	"log"
	"sync"
	"sync/atomic"
)

// use to manage go state
// usually use for counters

func main() {
	var ops atomic.Uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				ops.Add(1)
			}

			wg.Done()
		}()
	}

	wg.Wait()
	log.Println("counter:", ops.Load())
}
