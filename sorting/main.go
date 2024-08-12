package main

import (
	"log"
	"slices"
)

func main() {
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	log.Println(strs)

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	log.Println(ints)

	s := slices.IsSorted(strs)
	log.Println(s)
}
