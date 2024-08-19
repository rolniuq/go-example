package main

import "fmt"

// interface are named collections of method signatures

// Under the hood, interface values can be thought of as a tuple of a value and a concrete type:
// (value, type)
// An interface value holds a value of a specific underlying concrete type.
// Calling a method on an interface value executes the method of the same name on its underlying type.

type geometry interface {
	area() float64
	perim() float64
}

type myerr struct {
}

func (myerr) Error() string {
	return "myerr"
}

func checkNil(err error) bool {
	return err == nil
}

func main() {
	var myerr *myerr
	var err error = myerr
	fmt.Println(checkNil(err)) // false
}
