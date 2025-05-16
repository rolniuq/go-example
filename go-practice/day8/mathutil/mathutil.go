package mathutil

import "fmt"

func Sum(a, b int) int {
	return a + b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide 0")
	}

	return a / b, nil
}

func ProcessItem(items []int, fn func(int) int) []int {
	for i := 0; i < len(items); i++ {
		items[i] = fn(i)
	}

	return items
}
