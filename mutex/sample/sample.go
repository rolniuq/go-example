package sample

import (
	"fmt"
	"sync"
)

type Account struct {
	Balance int
	Name    string
}

func (a *Account) Deposit(amount int) {
	a.Balance += amount
}

func (a *Account) GetBalance() int {
	return a.Balance
}

func (a *Account) Withdraw(amount int) {
	a.Balance -= amount
}

// run: go run -race main.go
func RaceConditionSample() string {
	var account Account
	account.Name = "John"

	wg := sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			wg.Done()
			account.Deposit(100)
		}()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			wg.Done()
			account.Withdraw(100)
		}()
	}

	wg.Wait()
	return fmt.Sprintf("Wrong sample %d", account.GetBalance())
}
