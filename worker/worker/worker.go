package worker

type Worker[T any] struct {
}

func NewWorker[T any]() *Worker[T] {
	return &Worker[T]{}
}
