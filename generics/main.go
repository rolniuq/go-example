package main

import "fmt"

func MapKey[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}

	return r
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elements []T

	for e := lst.head; e != nil; e = e.next {
		elements = append(elements, e.val)
	}

	return elements
}

func main() {
	m := map[int]string{1: "2", 2: "2", 3: "3"}
	fmt.Println(MapKey(m))

	_ = MapKey[int, string](m)
	lst := List[int]{}
	lst.Push(10)
	lst.Push(23)
	fmt.Println(lst.GetAll())
}
