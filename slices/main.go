package main

import "fmt"

// Slice is supper powerful in GO
// want to know more? practice more :v

func main() {
	s := make([]int, 0)
	for i := 0; i < 3; i++ {
		s = append(s, i)
	}
	fmt.Println(s)
}
