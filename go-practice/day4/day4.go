package day4

import "fmt"

type Operation func(a, b int) int

func Apply(op Operation, a, b int) int {
	return op(a, b)
}

func MakeCounter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

type DivideError struct {
	Op   string
	A, B int
}

func (de *DivideError) Error() string {
	return fmt.Sprintf("cannot divide %d by %d", de.A, de.B)
}

func SafeDiv(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivideError{"div", a, b}
	}

	return a / b, nil
}

func WrapPanic(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("wrapped panic: %v", r)
		}
	}()

	fn()

	return
}

type Day4 struct{}

func (d *Day4) Exec() {
	fmt.Println(Apply(func(a, b int) int { return a + b }, 1, 2))
	fmt.Println(Apply(func(a, b int) int { return a - b }, 1, 2))

	counter := MakeCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	if _, err := SafeDiv(1, 0); err != nil {
		fmt.Println(err)
	}

	if err := WrapPanic(func() {
		panic("panic")
	}); err != nil {
		fmt.Println(err)
	}
}
