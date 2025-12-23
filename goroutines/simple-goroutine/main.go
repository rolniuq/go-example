package main

type Status int

const (
	PANICKED Status = iota
	RUNNING
)

type Goroutine struct {
	state      Status
	panicValue any
	deferFuncs []func()
}

func NewGoroutine() *Goroutine {
	return &Goroutine{}
}

func (g *Goroutine) AddDefer(f func()) {
	g.deferFuncs = append(g.deferFuncs, f)
}

func (g *Goroutine) MockPanic(value any) {
	g.state = PANICKED
	g.panicValue = value

	for i := len(g.deferFuncs) - 1; i >= 0; i-- {
		g.deferFuncs[i]()
		if g.state != PANICKED {
			return
		}
	}
}

func (g *Goroutine) MockRecover() any {
	if g.state == PANICKED {
		v := g.panicValue
		g.state = RUNNING
		g.panicValue = nil

		return v
	}

	return nil
}

func main() {
	g := NewGoroutine()

	g.AddDefer(func() {
		println("defer 1")
		if r := g.MockRecover(); r != nil {
			println("mock recover defer 1:", r.(string))
		}
	})

	g.AddDefer(func() {
		println("defer 2")
		if r := g.MockRecover(); r != nil {
			println("mock recover defer 2:", r.(string))
		}
	})

	g.MockPanic("panic value")
}
