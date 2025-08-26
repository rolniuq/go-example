package reader

import (
	"bufio"
	"os"
)

func OldReadLines(fileName string) ([]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func NewReadLines(fileName string) func(func(string) bool) {
	return func(yield func(string) bool) {
		f, err := os.Open(fileName)
		if err != nil {
			return
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if !yield(scanner.Text()) {
				return
			}
		}
	}
}
