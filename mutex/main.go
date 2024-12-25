package main

import (
	"log"
	"mutex/sample"
	"sync"
)

// mutex: mutual exclusion to avoid race condition
// ensure that only one goroutine at a time can access a resource

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.counters[name]++
}

func Mutex() {
	c := Container{
		counters: map[string]int{
			"a": 0,
			"b": 0,
		},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, times int) {
		for i := 0; i < times; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(4)
	go doIncrement("a", 100)
	go doIncrement("a", 100)
	go doIncrement("a", 100)
	go doIncrement("b", 100)

	wg.Wait()
	log.Println(c.counters)
}

func main() {
	// Mutex()
	sample.RaceConditionSample()
}
