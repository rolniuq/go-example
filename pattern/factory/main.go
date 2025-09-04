package main

import "fmt"

type BurgerFactory struct {
	name  string
	sauce string
	size  string
}

func newBurgerFactory() *BurgerFactory {
	return &BurgerFactory{}
}

func (b *BurgerFactory) createCheeseBurger() *BurgerFactory {
	b.name = "cheese burger"
	b.sauce = "cheese"
	b.size = "M"

	return b
}

func (b *BurgerFactory) createVeganBurger() *BurgerFactory {
	b.name = "vegan"
	b.sauce = "vegan"
	b.size = "X"

	return b
}

func main() {
	burgerFactory := newBurgerFactory()

	cheeseBurger := burgerFactory.createCheeseBurger()
	fmt.Println(cheeseBurger.name)

	veganBurger := burgerFactory.createVeganBurger()
	fmt.Println(veganBurger.name)
}
