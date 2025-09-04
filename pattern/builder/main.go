package main

import "fmt"

type BurgerBuilder struct {
	name  string
	sauce string
	size  string
}

func newBuggerBuilder() *BurgerBuilder {
	return &BurgerBuilder{}
}

func (b *BurgerBuilder) addName(name string) *BurgerBuilder {
	b.name = name

	return b
}

func (b *BurgerBuilder) addSauce(sauce string) *BurgerBuilder {
	b.sauce = sauce

	return b
}

func (b *BurgerBuilder) addSize(size string) *BurgerBuilder {
	b.size = size

	return b
}

func (b *BurgerBuilder) build() *BurgerBuilder {
	return b
}

func main() {
	burgerBuilder := newBuggerBuilder()
	burger := burgerBuilder.addName("cheese").addSauce("ot").addSize("L").build()
	fmt.Println(burger.name)
}
