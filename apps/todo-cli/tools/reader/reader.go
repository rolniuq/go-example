package reader

import (
	"fmt"
	"os"
	"todo-cli/tools/creator"
)

type Reader struct {
	readerPath string
}

func NewReader(readerPath string) *Reader {
	if err := creator.CreateFile(readerPath); err != nil {
		fmt.Println("error when create reader path", err)
	}

	return &Reader{
		readerPath: readerPath,
	}
}

func (r *Reader) Read() ([]byte, error) {
	content, err := os.ReadFile(r.readerPath)
	if err != nil {
		return nil, err
	}
	if len(content) == 0 {
		return nil, nil
	}

	return content, nil
}
