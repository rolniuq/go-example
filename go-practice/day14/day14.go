package day14

import "go-practice/day14/server"

type Day14 struct{}

func (d *Day14) Exec() {
	server := server.Server{}
	server.Exec()
}
