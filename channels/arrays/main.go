package main

import "fmt"

func main() {
	ch := make(chan int)

	arr := [3]int{1, 2, 3}
	for v := range arr {
		go func(val int) {
			ch <- val
		}(v)
	}

	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("DONE")
}
