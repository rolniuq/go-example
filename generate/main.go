package main

import "fmt"

type Person struct {
	Name string
}

//go:generate stringer -type=MyEnum
type MyEnum int

//go:generate stringer -type=MyPerson
type MyPerson uint

type MyType[T comparable] struct {
	Value T
}

const (
	First MyEnum = iota
	Second
	Third

	PersonA MyPerson = iota
)

func main() {
	fmt.Println(Second.String())
}
