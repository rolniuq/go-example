package main

import "fmt"

func main() {
	done := make(chan bool)
	unBuf := make(chan int)

	defer func() {
		close(unBuf)
		close(done)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			unBuf <- i
		}

		done <- true
	}()

	end := false
	for !end {
		select {
		case v := <-unBuf:
			fmt.Println(v)
		case <-done:
			fmt.Println("return")
			end = true

			break // break out of the select -> DONE will be printed
			// return // return immediately. DONE will not print
		}
	}

	fmt.Println("DONE")
}
