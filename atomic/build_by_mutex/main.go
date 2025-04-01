package main

import (
	"sync"
)

type AtomicInt64 struct {
	mu  sync.Mutex
	val int64
}

func (a *AtomicInt64) Add(n int64) int64 {
	a.mu.Lock()

	a.val += n
	newVal := a.val

	a.mu.Unlock()

	return newVal
}

func (a *AtomicInt64) Load() int64 {
	a.mu.Lock()

	val := a.val

	a.mu.Unlock()

	return val
}

func main() {
	var counter AtomicInt64

	var wg sync.WaitGroup

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for range 10 {
				counter.Add(1)
			}

		}()
	}

	wg.Wait()

	println(counter.Load())
}
