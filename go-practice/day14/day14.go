package day14

import (
	"go-practice/day14/client"
	"go-practice/day14/server"
)

type Day14 struct{}

func (d *Day14) Exec() {
	server := server.Server{}
	go func() {
		server.Exec()
	}()

	client := client.Client{}
	go func() {
		client.Exec()
	}()

	select {}
}
