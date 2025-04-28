package main

import "go-practice/day1"

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
		Register(&day1.Day1{})

	for _, d := range ds.ds {
		d.Exec()
	}
}
