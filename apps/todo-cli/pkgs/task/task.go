package task

import (
	"errors"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

func NewTask(title *string) (*Task, error) {
	t := Task{}
	if err := t.validate(title); err != nil {
		return nil, err
	}

	t.Title = *title

	return &t, nil
}

func (t *Task) validate(title *string) error {
	if title == nil {
		return errors.New("title cannot be empty")
	}

	return nil
}

func (t *Task) MarkCompleted() {
	t.Completed = true
}

func (t *Task) MarkIncomplete() {
	t.Completed = false
}
