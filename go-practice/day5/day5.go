package day5

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const times = 5
const max = 20

func generateNumber() <-chan int {
	ch := make(chan int, times)

	for i := 0; i < times; i++ {
		ch <- rand.Intn(max)
	}
	close(ch)

	return ch
}

func Merge(chans ...<-chan int) <-chan int {
	ch := make(chan int, len(chans)*times)

	wg := sync.WaitGroup{}
	wg.Add(len(chans))

	for _, c := range chans {
		go func(c <-chan int) {
			defer wg.Done()
			for n := range c {
				ch <- n
			}
		}(c)
	}

	wg.Wait()

	close(ch)

	return ch
}

func FanInFanOut() {
	ch1 := make(chan int, times)
	ch2 := make(chan int, times)
	ch3 := make(chan int, times)

	go func() {
		for n := range generateNumber() {
			ch1 <- n
		}

		close(ch1)
	}()

	go func() {
		for n := range generateNumber() {
			ch2 <- n
		}

		close(ch2)
	}()

	go func() {
		for n := range generateNumber() {
			ch3 <- n
		}

		close(ch3)
	}()

	ch := Merge(ch1, ch2, ch3)
	for c := range ch {
		fmt.Println(c)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- j * j
	}
}

func WorkerPool() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	for w := 1; w <= 3; w++ {
		go worker(jobs, results)
	}

	for i := 0; i < 10; i++ {
		jobs <- i
	}
	close(jobs)

	for a := 0; a < 10; a++ {
		fmt.Println(<-results)
	}
	close(results)
}

func Pipeline(ctx context.Context) {
	chan1 := make(chan int, 5)
	chan2 := make(chan int, 5)
	chan3 := make(chan int, 5)

	go func() {
		defer close(chan1)
		for i := 1; i <= 5; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(200 * time.Millisecond)
				chan1 <- i
			}
		}
	}()

	go func() {
		defer close(chan2)
		for c := range chan1 {
			select {
			case <-ctx.Done():
				return
			default:
				chan2 <- c * 2
			}
		}
	}()

	go func() {
		defer close(chan3)
		for c := range chan2 {
			select {
			case <-ctx.Done():
				return
			default:
				chan3 <- c + 1
			}
		}
	}()

	for c := range chan3 {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(c)
		}
	}
}

type Day5 struct{}

func (d *Day5) Exec() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	FanInFanOut()
	WorkerPool()
	Pipeline(ctx)
}
