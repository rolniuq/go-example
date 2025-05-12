package day6

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)

type ctxKey string

const TraceKey ctxKey = "traceID"
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

	ctx = d.Trace(ctx)

	ctx1, cancel1 := context.WithTimeout(ctx, 800*time.Millisecond)
	defer cancel1()

	ctx2, cancel2 := context.WithTimeout(ctx1, 500*time.Millisecond)
	defer cancel2()

	ctx3, cancel3 := context.WithTimeout(ctx2, 300*time.Millisecond)
	defer cancel3()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-time.After(900 * time.Millisecond):
			fmt.Println("[traceID]:", ctx.Value(TraceKey), "global timeout")
		case <-ctx3.Done():
			fmt.Printf("[traceID]:", ctx.Value(TraceKey), "ctx3 done: %v\n", ctx3.Err())
		case <-ctx2.Done():
			fmt.Println("[traceID]:", ctx.Value(TraceKey), "ctx2 done")
		case <-ctx1.Done():
			fmt.Println("[traceID]:", ctx.Value(TraceKey), "ctx1 done")
		}
	}()
	wg.Wait()
}

func (d *Day6) Server() {
	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = d.Trace(ctx)

		for range 5 {
			select {
			case <-ctx.Done():
				w.WriteHeader(504)
				return
			default:
				time.Sleep(200 * time.Millisecond)
			}
		}

		w.Write([]byte("done"))
	})

	http.ListenAndServe(":3000", nil)
}

func (d *Day6) Trace(ctx context.Context) context.Context {
	traceId := ctx.Value(TraceKey)
	if traceId == nil {
		ctx = context.WithValue(ctx, TraceKey, uuid.New())
	}

	return ctx
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
	d.Server()
}
