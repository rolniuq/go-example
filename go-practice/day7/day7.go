package day7

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/singleflight"
)

type Day7 struct{}

func (d *Day7) ErrGroupExample() {
	ran := func() int {
		return rand.Intn(200)
	}

	eg := errgroup.Group{}

	for i := 0; i < 5; i++ {
		i := i

		eg.Go(func() error {
			sleep := ran()
			time.Sleep(time.Duration(sleep) * time.Millisecond)
			if sleep%2 == 0 {
				return fmt.Errorf("even error %d", i)
			}

			fmt.Printf("task %d done\n", i)

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Printf("errgroup failed: %v", err)
	} else {
		fmt.Println("errgroup all success")
	}
}

func (d *Day7) SemaphoreExample() {
	wg := sync.WaitGroup{}
	wg.Add(10)

	sem := make(chan struct{}, 3)
	for i := 0; i < 10; i++ {
		i := i

		sem <- struct{}{}
		go func(i int) {
			defer wg.Done()
			defer func() { <-sem }()
			fmt.Printf("task %d start \n", i)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("task %d done \n", i)
		}(i)
	}

	wg.Wait()
}

func (d *Day7) RateLimiterExample() {
	ticker := time.NewTicker(500 * time.Millisecond)
	for i := range 5 {
		<-ticker.C
		fmt.Printf("request %d at %v\n", i, time.Now())
	}

	ticker.Stop()
}

func (d *Day7) SingleflightExample() {
	wg := sync.WaitGroup{}
	var g singleflight.Group

	wg.Add(3)
	for range 3 {
		go func() {
			defer wg.Done()
			v, err, shared := g.Do("key1", func() (interface{}, error) {
				fmt.Println("actually fetching")
				time.Sleep(100 * time.Millisecond)
				return "result", nil
			})
			fmt.Println("got", v, "shared?", shared, "err", err)
		}()
	}

	wg.Wait()
}

type M struct {
	mu sync.Mutex
	n  int
}

func (m *M) Inc() { m.mu.Lock(); m.n++; m.mu.Unlock() }

func (d *Day7) CounterWithMutex() {
	m := &M{}

	for range 5 {
		m.Inc()
	}

	fmt.Println("counter with mutex", m.n)
}

func (d *Day7) CounterWithChannel() {
	incr := make(chan struct{})
	done := make(chan struct{})
	go func() {
		var n int
		for {
			select {
			case <-incr:
				n++
			case <-done:
				fmt.Println("count =", n)
				return
			}
		}
	}()
	// spawn 5 increments:
	for range 5 {
		incr <- struct{}{}
	}
	done <- struct{}{}
}

func (d *Day7) Exec() {
	d.ErrGroupExample()
	d.SemaphoreExample()
	d.RateLimiterExample()
	d.SingleflightExample()
	d.CounterWithMutex()
	d.CounterWithChannel()
}
