package mem

func StackAlloc(x int) *int {
	y := x
	return &y
}

func HeapAlloc(x int) *int {
	z := new(int)
	*z = x
	return z
}
