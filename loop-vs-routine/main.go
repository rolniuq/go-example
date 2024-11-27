package main

import (
	"fmt"
	"sync"
	"time"
)

func loop() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		fmt.Println("loop", i)
	}
	fmt.Println("loop", time.Since(start))
}

func routine() {
	start := time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("routine", i)
		}(i)
	}
	wg.Wait()

	fmt.Println("routine", time.Since(start))
}

func main() {
	loop()
	routine()
}
