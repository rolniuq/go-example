package mathutil

import "fmt"

func Sum(a, b int) int {
	return a + b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %d by %d", a, b)
	}

	return a / b, nil
}

func ProcessItems(items []int, fn func(int) int) []int {
	for i := 0; i < len(items); i++ {
		items[i] = fn(items[i])
	}

	return items
}

type ItemProcessor interface {
	Process(int) int
}
