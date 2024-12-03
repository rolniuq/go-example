package main

import "fmt"

/*
a label is a way to mark a specific line of code in your program
typically used in conjunction with control-flow statements: `goto``, `break``, or `continue`.
Labels are identifiers followed by a colon (:)
*/

func loop() {
outerLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("i: %d, j: %d\n", i, j)
			if i == 1 && j == 1 {
				break outerLoop // Break out of the outer loop
			}
		}

	}
	fmt.Println("Exited loops")
}

func loop2() {
outerLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				continue outerLoop // Skip the rest of the outer loop iteration
			}
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}

	fmt.Println("Exited loops")
}

func skip() {
	isEnd := true

	fmt.Println(1)
	if isEnd {
		goto End
	}

	fmt.Println(2)

End:
	fmt.Println(3)
}

func main() {
	skip()
	loop()
	loop2()
}
