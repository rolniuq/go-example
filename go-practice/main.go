package main

import (
	"go-practice/day14"
)

type Daily interface {
	Exec()
}

type Dailies struct {
	ds []Daily
}

func NewDailies() *Dailies {
	return &Dailies{}
}

func (ds *Dailies) Register(d Daily) *Dailies {
	ds.ds = append(ds.ds, d)

	return ds
}

func main() {
	ds := NewDailies().
		//Register(&day1.Day1{}).
		//Register(&day2.Day2{}).
		//Register(&day3.Day3{}).
		//Register(&day4.Day4{}).
		//Register(&day5.Day5{}).
		//Register(&day6.Day6{}).
		//Register(&day7.Day7{}).
		//Register(&day9.Day9{}).
		//Register(&day10.Day10{}).
		//Register(&day11.Day11{}).
		//Register(&day13.Day13{}).
		Register(&day14.Day14{})

	for _, d := range ds.ds {
		d.Exec()
	}
}
