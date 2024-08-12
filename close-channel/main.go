package main

import "fmt"

/*
close channel indicate that there is no value send to this channel
syntax: close(channel)

-> very useful to communicate completion to the channel's receivers
*/
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("receive job", j)
			} else {
				fmt.Println("receive all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 5; j++ {
		jobs <- j
		fmt.Println("send jobs")
	}
	close(jobs)
	fmt.Println("close jobs")

	<-done
	_, ok := <-jobs
	fmt.Println("more jobs", ok)
}
