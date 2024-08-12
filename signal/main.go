package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// we need to graceful shutdowns when system receive SIGTERM or command line tool to stop process input if system receive SIGINT
// this example is how to do

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("waiting signal")
	<-done
	fmt.Println("exiting")
}
