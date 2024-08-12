package main

import "fmt"

// go support method of struct type of struct

type rect struct {
	w int
	h int
}

func (r *rect) area() int {
	return r.h * r.w
}

func (r *rect) perim() int {
	return 2*r.w + 2*r.h
}

func main() {
	r := rect{w: 10, h: 10}

	fmt.Println(r.perim())
	fmt.Println(r.area())
}
