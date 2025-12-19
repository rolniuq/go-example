package main

import (
	"fmt"
	"sync"
	"time"
)

type semaphore struct {
	mu   sync.Mutex
	cond sync.Cond

	count int
}

func NewSemaphore(n int) *semaphore {
	s := &semaphore{count: n}
	s.cond = sync.Cond{L: &s.mu}
	return s
}

// take a slot in semaphore
func (s *semaphore) Accquire() {
	s.mu.Lock()

	if s.count == 0 {
		// force the goroutine to wait
		s.cond.Wait()
	}
	s.count--

	s.mu.Unlock()
}

func (s *semaphore) Release() {
	s.mu.Lock()

	s.count++
	// signial to wake up a waiting goroutine
	s.cond.Signal()

	s.mu.Unlock()
}

func worker(i int, sem *semaphore) {
	fmt.Println("Worker is waiting for slot", i)
	sem.Accquire()
	fmt.Println("Worker is using slot", i)
	time.Sleep(500 * time.Millisecond)
	sem.Release()
	fmt.Println("Worker released slot", i)
}

func main() {
	sem := NewSemaphore(3)
	for i := range 20 {
		go worker(i, sem)
	}

	time.Sleep(5 * time.Second)
}
