package worker

import "fmt"

type Work interface {
	Do(id int, jobs <-chan int, results chan<- int, done <-chan bool)
}

type WorkerA struct{}

func (w WorkerA) Do(id int, jobs <-chan int, results chan<- int, done <-chan bool) {
	for {
		select {
		case job := <-jobs:
			fmt.Println("Worker", id, "started  job", job)
			results <- job * 2
			fmt.Println("Worker", id, "finished job", job)
		case <-done:
			fmt.Println("Worker", id, "done")
			return
		}
	}
}

type WorkerPool struct {
	numberWorkers int

	jobs    chan int
	results chan int
	done    chan bool

	works []Work
}

func NewWorkerPool(n int) *WorkerPool {
	return &WorkerPool{
		numberWorkers: n,

		jobs:    make(chan int),
		results: make(chan int),
		done:    make(chan bool, n),
	}
}

func (w *WorkerPool) Register(work Work) {
	w.works = append(w.works, work)
}

func (w *WorkerPool) Do() {
	w.jobs = make(chan int, len(w.works))

	for i := 0; i < len(w.works); i++ {
		go w.works[i].Do(i, w.jobs, w.results, w.done)
	}

	w.done <- true

	for res := range w.results {
		fmt.Println("Result: ", res)
	}

	close(w.jobs)
	close(w.results)
	close(w.done)
}
