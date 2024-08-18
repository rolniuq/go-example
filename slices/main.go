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

	str := "helloworld"
	split1 := str[5:] // world
	fmt.Println(split1)
	split2 := str[0:5] // hello
	fmt.Println(split2)

	// (start:end]
}
