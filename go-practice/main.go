package main

import (
	"go-practice/day7"
)

//"go-practice/day1"
//"go-practice/day2"
//"go-practice/day3"
//"go-practice/day4"
//"go-practice/day5"

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
		// Register(&day6.Day6{}).
		Register(&day7.Day7{})

	for _, d := range ds.ds {
		d.Exec()
	}
}
