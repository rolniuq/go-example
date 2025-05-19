package day9

import (
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
)

type Day9 struct{}

func (d *Day9) Fib(n int) int {
	if n <= 1 {
		return n
	}

	return d.Fib(n-1) + d.Fib(n-2)
}

func (d *Day9) RunCPUProfile(n int) {
	f, _ := os.Create("out/cpu.out")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for i := range n {
		_ = d.Fib(i)
	}
}

func (d *Day9) RunMemProfile() {
	f, _ := os.Create("out/mem.out")
	runtime.GC()
	pprof.WriteHeapProfile(f)
	f.Close()
}

func (d *Day9) Trace() {
	f, _ := os.Create("out/trace.out")
	trace.Start(f)
	defer trace.Stop()
	d.RunCPUProfile(10)
}

func (d *Day9) Exec() {
	d.RunCPUProfile(20)

	_ = make([]int, 1e6)
	d.RunMemProfile()
}
