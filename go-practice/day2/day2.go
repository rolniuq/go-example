package day2

import "fmt"

type Mover interface {
	Move() string
}

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d *Dog) Move() string {
	return "walk"
}

func (d *Dog) Speak() string {
	return "woof"
}

type RobotDog struct {
	Model string
}

func (d *RobotDog) Move() string {
	return "drive"
}

func PrintActions(m Mover, s Speaker) {
	fmt.Println(m.Move())
	fmt.Println(s.Speak())
}

type Day2 struct{}

func (d *Day2) Exec() {
	dog := &Dog{Name: "Fido"}
	PrintActions(dog, dog)

	robotDog := &RobotDog{Model: "R2D2"}
	fmt.Println(robotDog.Move())
}
