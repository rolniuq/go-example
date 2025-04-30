package main

func main() {
	// l := &Lesson1{}
	// l.Execute()
	// l.Fixed()

	l2 := NewWorkerPool(5, 3)
	l2.Execute()
}
