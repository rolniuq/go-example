package day10

import "fmt"

type Day10 struct {
}

func Map[T, U any](arr []T, f func(T) U) []U {
	if arr == nil {
		return nil
	}

	res := make([]U, len(arr))
	for i, v := range arr {
		res[i] = f(v)
	}

	return res
}

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var v T
		return v, false
	}

	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return v, true
}

type Number interface {
	~int | ~float64
}

func SumAll[T Number](items ...T) T {
	var sum T
	for _, v := range items {
		sum += v
	}

	return sum
}

func Filter[T any](items []T, fn func(T) bool) []T {
	var res []T
	for _, v := range items {
		if fn(v) {
			res = append(res, v)
		}
	}

	return res
}

func (d *Day10) Exec() {
	vs := []int{1, 2, 3}
	ss := Map(vs, func(v int) string { return fmt.Sprintf("V%d", v) })
	fmt.Println(ss)

	s := Stack[string]{}
	s.Push("a")
	s.Push("b")
	v, ok := s.Pop()
	fmt.Println(v, ok)

	fmt.Println(SumAll(1, 2, 3))
	fmt.Println(SumAll(1.5, 2.5, 3.5))

	nums := []int{1, 2, 3, 4, 5}
	evens := Filter(nums, func(v int) bool { return v%2 == 0 })
	fmt.Println(evens)
}
