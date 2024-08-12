package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("%d", b.num)
}

type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "text",
	}

	fmt.Println(co.num, co.str)

	type d interface {
		describe() string
	}

	var d1 d = co
	fmt.Println(d1.describe())
}
