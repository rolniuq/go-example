package pointer

func NewPointer[T any](v T) *T {
	return &v
}

func PointerValue[T any](v *T) T {
	var t T
	if v != nil {
		t = *v
	}

	return t
}
