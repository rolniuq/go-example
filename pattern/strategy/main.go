package main

import "strategy/custom"

// The Strategy pattern is a behavioral design pattern that allows you to define a family of algorithms, encapsulate each one as a separate object, and make them interchangeable.
// The Strategy pattern enables you to select an algorithm's behavior at runtime

type Strategy interface {
	Sort([]int)
}

type Context struct {
	strategy Strategy
}

func (c *Context) setStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) execStrategy(nums []int) {
	c.strategy.Sort(nums)
}

func main() {
	data := []int{34, 7, 23, 32, 5, 62}
	c := Context{}

	// with bubble
	bubble := custom.BubbleSort{}
	c.setStrategy(&bubble)
	c.execStrategy(data)

	// with quick
	quick := custom.QuickSort{}
	c.setStrategy(&quick)
	c.execStrategy(data)
}
