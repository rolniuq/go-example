package main

import "fmt"

// function is a cental in GO -> cuz there is no class, only function

func basicFunc(a, b int) int {
	return a + b
}

func funcMulReturns() (int, int) {
	return 1, 2
}

type x complex64

// variadic functions meaning a function biến thiên
func variadicFunc(nums ...int) int {
	total := 0
	for _, num := range nums {
		total = total + num
	}

	return total
}

func main() {
	b := basicFunc(1, 2)
	fmt.Println(b)
	funcMulReturns()

	// can call like this
	variadicFunc(1, 2, 3)
	// and like this also
	nums := []int{1, 2, 3}
	variadicFunc(nums...)
}
