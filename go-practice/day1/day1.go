package day1

import "fmt"

type user struct {
	name string
	age  int64
}

type account struct {
	owner  *user
	amount float64
}

func (a *account) deposit(amount float64) {
	defer fmt.Println("Deposit finished.")
	if amount <= 0 {
		panic("Invalid deposit amount")
	}

	a.amount += amount
}

type Day1 struct{}

func (d *Day1) Exec() {
	user := &user{
		name: "quynh",
		age:  25,
	}

	account := &account{
		owner:  user,
		amount: 2000.0,
	}

	account.deposit(10.0)

	fmt.Println("Day1:", account.amount)
}
