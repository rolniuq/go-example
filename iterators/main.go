// Iterators as iter. Valid from 1.23
package main

import (
	"fmt"
	"iter"
)

type element[T any] struct {
	val  T
	next *element[T]
}

type List[T any] struct {
	head, tail *element[T]
}

func New[T any]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Push(val T) {
	if l.tail == nil {
		head := &element[T]{val: val}
		l.head = head
		l.tail = head
	} else {
		l.tail.next = &element[T]{val: val}
		l.tail = l.tail.next
	}
}

func (l *List[T]) All() iter.Seq[T] {
	return func(yield func(V T) bool) {
		for e := l.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func genFib() iter.Seq[int] {
	return func(yield func(V int) bool) {
		a, b := 0, 1
		for {
			if !yield(a) {
				return
			}

			a, b = b, a+b
		}
	}
}

type country struct {
	languages []string
}

func NewCountry(languages []string) *country {
	return &country{languages}
}

func (c *country) PrintLanguages() iter.Seq[string] {
	return func(yield func(V string) bool) {
		for _, l := range c.languages {
			if !yield(l) {
				return
			}
		}
	}
}

func looper[T any](a []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range a {
			if !yield(v) {
				return
			}
		}
	}
}

func main() {
	c := NewCountry([]string{"English", "Spanish", "French"})
	for l := range c.PrintLanguages() {
		fmt.Println(l)
	}

	l := New[int]()
	l.Push(1)
	l.Push(2)
	l.Push(3)
	for v := range l.All() {
		fmt.Println(v)
	}

	for v := range genFib() {
		if v > 10 {
			break
		}

		fmt.Println(v)
	}

	for v := range looper([]string{"a", "b", "c"}) {
		fmt.Println(v)
	}

	for v := range looper([]int{1, 2, 3, 4, 5}) {
		fmt.Println(v)
	}

	type person struct {
		name string
		age  int
	}

	people := []person{{"Alice", 20}, {"Bob", 30}, {"Charlie", 40}}
	for v := range looper(people) {
		fmt.Println(v)
	}
}
