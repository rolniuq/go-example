package main

import (
	"fmt"
	"time"
)

// which we will run several concurrent instances
// these worker will receive work on the jobs channel and send the corresponding result on the results channel
// we will sleep a second per job to simulate an expensive task
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	// in order to use or pool of workers we need to send them work and collect their results
	// we make 2 channels for this
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// this starts up 3 workers, initial blocked because there are no jobs yet
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// we send 5 jobs and then close that channel to indicate thats all the work we have
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	// we collect all the results of the work
	// this also ensure that the worker goroutine have finished
	// wait for multiple goroutines is to use a wait group
	for i := 1; i <= numJobs; i++ {
		<-results
	}
}
