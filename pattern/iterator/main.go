package main

import "fmt"

// The Iterator pattern is a design pattern that provides a way to access the elements of a collection (like arrays or lists)
// sequentially without exposing its underlying structure.

type Collection struct {
	items []int
}

type Iterator interface {
	hasNext() bool
	next() int
}

type CollectionIterator struct {
	collection *Collection
	index      int
}

func (c *CollectionIterator) hasNext() bool {
	return c.index < len(c.collection.items)
}

func (c *CollectionIterator) next() int {
	if !c.hasNext() {
		return -1
	}

	res := c.collection.items[c.index]
	c.index++

	return res
}

func newCollectionIterator(c *Collection) Iterator {
	return &CollectionIterator{
		collection: c,
		index:      0,
	}
}

func main() {
	c := &Collection{
		items: []int{1, 2, 3, 4, 5, 6},
	}

	iterator := newCollectionIterator(c)
	for iterator.hasNext() {
		fmt.Println(iterator.next())
	}
}
