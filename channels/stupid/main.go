package main

import "fmt"

// deadlock chan

func main() {
	ch := make(chan int)
	ch <- 1

	select {
	case i := <-ch:
		fmt.Println(i)
		return
	default:
		fmt.Println("default")
		return
	}
}
