package main

import (
	"fmt"
	"sync"
	"time"
)

var mtx sync.Mutex
var cond = sync.NewCond(&mtx)

func dummyGoroutine(id int) {
	cond.L.Lock()
	defer cond.L.Unlock()

	fmt.Println("go routine is waiting", id)
	cond.Wait()

	fmt.Println("go routine receive signal", id)
}

func main() {
	go dummyGoroutine(1)

	time.Sleep(1 * time.Second)

	fmt.Println("send signal")
	cond.Signal()

	time.Sleep(1 * time.Second)
}
