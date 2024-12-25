package comparabletest

type Person[T comparable] struct {
	Name  string
	Age   int
	Bonus T
}

func NewPerson[T comparable]() *Person[T] {
	return &Person[T]{}
}

func (p *Person[T]) SetName(name string) {
	p.Name = name
}

func (p *Person[T]) SetAge(age int) {
	p.Age = age
}

func (p *Person[T]) SetBonus(bonus T) {
	p.Bonus = bonus
}

func (p *Person[T]) GetName() string {
	return p.Name
}

func (p *Person[T]) GetAge() int {
	return p.Age
}

func (p *Person[T]) GetBonus() T {
	return p.Bonus
}

func (p *Person[T]) IsAdult() bool {
	return p.Age >= 18
}
