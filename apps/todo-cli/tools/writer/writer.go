package writer

import (
	"fmt"
	"os"
	"todo-cli/tools/creator"
)

type Writer struct {
	writerPath string
}

func NewWriter(writerPath string) *Writer {
	if err := creator.CreateFile(writerPath); err != nil {
		fmt.Println("error when create writer path", err)
	}

	return &Writer{
		writerPath: writerPath,
	}
}

func (w *Writer) Write(content string) error {
	file, err := os.Create(w.writerPath)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.WriteString(content); err != nil {
		return err
	}

	return nil
}
