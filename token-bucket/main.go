package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"time"
	"weak"

	"github.com/redis/go-redis/v9"
)

//go:embed script.lua
var script []byte

type tokenBucket struct {
	rbd        *redis.Client
	sc         *redis.Script
	capacity   int
	refillRate float64 // token per second
}

func NewTokenBucket(rbd *redis.Client, capacity int, refillRate float64) *tokenBucket {
	return &tokenBucket{
		rbd: rbd,
		sc:  redis.NewScript(string(script)),
	}
}

func (tb *tokenBucket) Allow(ctx context.Context, key string, cost int) (allowed bool, remaining float64, err error) {
	now := time.Now().UnixMilli()
	refillPerMs := tb.refillRate / 1000.0

	res, err := tb.sc.Run(
		ctx,
		tb.rbd,
		[]string{key},
		tb.capacity,
		refillPerMs,
		now,
		cost,
	).Result()

	if err != nil {
		return false, 0, err
	}

	arr := res.([]interface{})
	allowed = arr[0].(int64) == 1
	remaining = arr[1].(float64)

	return
}

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	tb := NewTokenBucket(
		rdb,
		10, // capacity
		5,  // refill 5 tokens / second
	)

	key := "token_bucket:user_123"

	for i := 1; i <= 15; i++ {
		allowed, remaining, err := tb.Allow(ctx, key, 1)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(
			"Request %02d | allowed=%v | remaining=%.2f\n",
			i,
			allowed,
			remaining,
		)

		time.Sleep(100 * time.Millisecond)
	}
}
