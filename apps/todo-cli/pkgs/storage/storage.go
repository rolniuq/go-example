package storage

import (
	"todo-cli/pkgs/task"
)

type Storage struct {
	data []task.Task
}

func NewStorage[T task.Task]() *Storage {
	return &Storage{}
}

func (s *Storage) Add(value task.Task) {
	if s == nil {
		return
	}

	if s.data == nil {
		s.data = make([]task.Task, 0)
	}

	s.data = append(s.data, value)
}

func (s *Storage) Remove(title string) {
	if s == nil {
		return
	}

	for i, v := range s.data {
		if v.Title == title {
			s.data = append(s.data[:i], s.data[i+1:]...)
			return
		}
	}
}

func (s *Storage) Edit(value task.Task) {
	if s == nil {
		return
	}

	for i, v := range s.data {
		if v.ID == value.ID {
			s.data[i] = value
			return
		}
	}
}

func (s *Storage) GetData() []task.Task {
	if s == nil {
		return nil
	}

	return s.data
}
