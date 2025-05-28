package main

import (
	"fmt"
	"time"
)

func main() {
	var data [][]byte
	for i := 0; i < 100000; i++ {
		buf := make([]byte, 1024)
		data = append(data, buf)
	}
	fmt.Println("Done creating")
	time.Sleep(5 * time.Second)
}
