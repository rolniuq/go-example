package main

import "sync"

var counter int

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	counter++
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	wg.Wait()
	println(counter)
}
