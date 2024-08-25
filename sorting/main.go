package main

import (
	"log"
	"slices"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	log.Println(strs)

	str := "acb"
	strRunes := []rune(str)
	sort.Slice(strRunes, func(i, j int) bool {
		return strRunes[i] < strRunes[j]
	})
	log.Println(str)

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	log.Println(ints)

	s := slices.IsSorted(strs)
	log.Println(s)
}
