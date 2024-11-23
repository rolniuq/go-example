package main

import (
	"fmt"
	"log"
)

func Fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * Fact(n-1)
}

func main() {
	var a int
	fmt.Println("enter a number:")
	if _, err := fmt.Scan(&a); err != nil {
		log.Fatal("failed to input a number: %v", err)
	}

	fmt.Println(Fact(a))
}
