package day14

import (
	"fmt"
	"go-practice/day14/client"
	"go-practice/day14/server"
	"sync"
)

type Day14 struct{}

func (d *Day14) Exec() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	server := server.Server{}
	go func() {
		defer wg.Done()
		server.Exec()
	}()

	client := client.Client{}
	go func() {
		defer wg.Done()
		client.Exec()
	}()

	wg.Wait()

	fmt.Println("Setup completed")

	select {}
}
