package main

import "fmt"

func basicSwitch() {
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("undefinded")
	}
}

func switchToDetectType(v any) {
	canUGuess := func() {
		switch t := v.(type) {
		case bool:
			fmt.Println("boolean")
		case int:
			fmt.Println("integer")
		case string:
			fmt.Println("string")
		default:
			fmt.Println("cannot detect type", t)
		}
	}

	canUGuess()
}

func main() {
	basicSwitch()
	switchToDetectType("String")
	switchToDetectType(3)
	switchToDetectType(false)
}
