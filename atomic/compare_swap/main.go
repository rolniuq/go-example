package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicAddInt64(addr *int64, delta int64) int64 {
	for {
		v := atomic.LoadInt64(addr) // load current value
		newVal := v + delta         // calculate new value

		if atomic.CompareAndSwapInt64(addr, v, newVal) {
			return newVal // return new value
		}
	}
}

func main() {
	var counter int64
	var wg sync.WaitGroup

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 10 {
				counter = AtomicAddInt64(&counter, 1)
			}
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
