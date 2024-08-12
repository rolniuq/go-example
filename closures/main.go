package main

import "fmt"

// basicly closure function is the function contain the other function
// the function inside can use variable out of itself scope
func closuresFunc() func() int {
	h := 1
	return func() int {
		h += 1
		return h
	}
}

func main() {
	// init closure func
	f1 := closuresFunc()
	fmt.Println(f1())
	f2 := f1()
	fmt.Println(f2)
	f3 := f1()
	fmt.Println(f3)

	// init again
	h1 := closuresFunc()
	fmt.Println(h1())
}
