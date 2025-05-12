package day6

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type ctxKey string

const UserKey ctxKey = "user"

type Day6 struct{}

func (d *Day6) ContextTree(ctx context.Context) {
	ctx1, cancel1 := context.WithDeadline(ctx, time.Now().Add(300*time.Millisecond))
	ctx2, cancel2 := context.WithTimeout(ctx1, 200*time.Millisecond)
	defer cancel1()
	defer cancel2()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-time.After(500 * time.Millisecond):
			fmt.Println("timeout waiting")
		case <-ctx2.Done():
			fmt.Printf("ctx2 cancelled: %v\n", ctx2.Err())
		}
	}()
	wg.Wait()
}

func (d *Day6) Fetch(ctx context.Context, url string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(400 * time.Millisecond):
		fmt.Println("fetched", url)
		return nil
	}
}

func (d *Day6) Process(ctx context.Context) {
	user := ctx.Value(UserKey)
	if user == nil {
		fmt.Println("Hello, guest!")
		return
	}

	fmt.Printf("Hello, %s!\n", user.(string))
}

func (d *Day6) ContextTreeV2() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer cancel()

	ctx1, cancel1 := context.WithTimeout(ctx, 800*time.Millisecond)
	defer cancel1()

	ctx2, cancel2 := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel2()

	ctx3, cancel3 := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel3()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx3.Done():
				fmt.Printf("ctx3 cancelled: %v\n", ctx3.Err())
				return
			case <-ctx2.Done():
				fmt.Printf("ctx2 cancelled: %v\n", ctx2.Err())
				return
			case <-ctx1.Done():
				fmt.Printf("ctx1 cancelled: %v\n", ctx1.Err())
				return
			case <-time.After(900 * time.Millisecond):
				fmt.Println("timeout waiting")
				return
			}
		}
	}()
	wg.Wait()
}

func (d *Day6) Exec() {
	// ctx := context.Background()
	// d.ContextTree(ctx)

	// ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	// defer cancel()
	// if err := d.Fetch(ctx, "http://example.com"); err != nil {
	// 	fmt.Println(err)
	// }

	// ctx = context.WithValue(ctx, UserKey, "Quynh")
	// d.Process(ctx)

	d.ContextTreeV2()
}
