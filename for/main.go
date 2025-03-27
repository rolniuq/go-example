package main

import "fmt"

func singleCondition() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
}

func classic() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}

func modern() {
	fmt.Println("For loop modern running")
	for i := range 5 {
		fmt.Println(i)
	}
	fmt.Println("For loop modern done")
}

func rangeLoop() {
	fmt.Println("range loop running")
	// with go version >= 1.22 we can iterate over a range with int
	for i := range [3]int{} {
		fmt.Println("range", i)
	}

	fmt.Println("range loop done")
}

func withoutCondition() {
	i := 1
	for {
		fmt.Println("loop without condition")
		if i == 3 {
			break
		}
		i = i + 1
	}
}

func withObj() {
	obj := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	for i := range obj {
		fmt.Println(i)
	}

	for _, v := range obj {
		fmt.Println(v)
	}
}

func main() {
	singleCondition()
	classic()
	rangeLoop()
	withoutCondition()
	withObj()
	modern()
}
