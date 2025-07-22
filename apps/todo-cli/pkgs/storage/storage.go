package storage

import (
	"fmt"
	"todo-cli/configs"
	"todo-cli/pkgs/task"
	"todo-cli/tools/converter"
	"todo-cli/tools/creator"
	"todo-cli/tools/reader"
	"todo-cli/tools/writer"
)

type Storage struct {
	data   []task.Task
	reader *reader.Reader
	writer *writer.Writer
}

func NewStorage[T task.Task](config *configs.Config) *Storage {
	if err := creator.CreateDir(config.StoragePath); err != nil {
		fmt.Println("error when create storage path", err)
	}

	storage := &Storage{
		reader: reader.NewReader(config.ReaderPath),
		writer: writer.NewWriter(config.WriterPath),
	}

	data, err := storage.reader.Read()
	if err != nil {
		fmt.Println("error when read storage path", err)
	}
	if data != nil {
		tasks, err := converter.ByteToType[[]task.Task](data)
		if err != nil {
			fmt.Println("error when convert string to type", err)
		}
		if tasks != nil {
			storage.data = *tasks
		}
	}

	return storage
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

func (s *Storage) Save() error {
	if s == nil {
		return nil
	}

	for i := range s.data {
		s.data[i].ID = i + 1
	}

	data, err := converter.TypeToString(s.data)
	if err != nil {
		return err
	}

	if err = s.writer.Write(data); err != nil {
		return err
	}

	return nil
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
