package main

import (
	"fmt"
	"sync"
)

func raiseError() error {
	return fmt.Errorf("raise error")
}

func working() error {
	select {}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() error {
		defer wg.Done()
		if err := raiseError(); err != nil {
			return err
		}

		return nil
	}()

	go func() error {
		defer wg.Done()
		if err := working(); err != nil {
			return err
		}

		return nil
	}()

	wg.Wait()

	fmt.Println("done")
}
