package main

import "fmt"

// hay còn có thể gọi là đệ quy
func fibonacci(n int) int {
	if n == 0 {
		return 1
	}

	return n + (n - 1)
}

func main() {
	res := fibonacci(3)
	fmt.Println(res)
}
