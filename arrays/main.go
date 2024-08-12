package main

import "fmt"

// array is NUMBERED sequence of elements of a specific length
// in go slice is more common and more useful

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	for _, v := range arr {
		fmt.Println(v)
	}

	fmt.Println(cap(arr))
}
