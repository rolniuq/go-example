package main

import (
	"fmt"
	"myerror/wrapper"
)

func main() {
	err := fmt.Errorf("this is an error")
	fmt.Println(err)
	println(err)

	err = wrapper.NewError(1, "type", "message", nil)
	fmt.Println(err)
	println(err)

	// cannot assign
	// err1 := wrapper.NewError(2, "type", "message", nil)
	// err1 = fmt.Errorf("this is an error: %w", err1)
}
