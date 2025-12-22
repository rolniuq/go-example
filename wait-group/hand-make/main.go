package main

import (
	"fmt"
	"sync"
	"time"
)

type WaitGroupHandMake struct {
	count int
	mu    sync.Mutex
	cond  sync.Cond
}

func NewWaitGroupHandMake() *WaitGroupHandMake {
	wg := &WaitGroupHandMake{}
	wg.cond = sync.Cond{L: &wg.mu}
	return wg
}

func (wg *WaitGroupHandMake) Add(delta int) {
	wg.mu.Lock()
	defer wg.mu.Unlock()

	wg.count += delta

	if wg.count < 0 {
		panic("negative WaitGroup counter")
	}

	if wg.count == 0 {
		wg.cond.Broadcast()
	}
}

func (wg *WaitGroupHandMake) Done() {
	wg.Add(-1)
}

func (wg *WaitGroupHandMake) Wait() {
	wg.mu.Lock()
	defer wg.mu.Unlock()

	for wg.count > 0 {
		wg.cond.Wait()
	}
}

func printWithTimeout() {
	fmt.Println("hello")
	time.Sleep(500 * time.Millisecond)
}

func main() {
	wg := NewWaitGroupHandMake()

	wg.Add(2)

	go func() {
		defer wg.Done()
		printWithTimeout()
	}()

	go func() {
		defer wg.Done()
		printWithTimeout()
	}()

	wg.Wait()
}
