package main

import (
	"fmt"
)

// how to define a go channel to receive a value
// p chan <- typeof value

// how to define a go channel to send a value
// p <- chan typeof value

// to run a go routine function
// use go keywork and the function
// ex: go fn()

func ping(c chan<- string, msg string) {
	c <- msg
}

func pong(p <-chan string, c chan<- string) {
	msg := <-p
	c <- msg
}

func f(from string) {
	for i := 0; i < 5; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// f("direct")

	// go func(s string) {
	// 	fmt.Println(s)
	// }("goo")

	// go f("gorountine")

	// go func(msg string) {
	// 	fmt.Println(msg)
	// }("going")

	// time.Sleep(time.Second)
	// fmt.Println("done")

	c := make(chan string, 1)
	p := make(chan string, 1)
	ping(c, "hello")
	pong(c, p)
	fmt.Println(<-p)
}
