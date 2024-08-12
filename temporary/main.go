package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// throughout program execution, we often want to create data that isn't needed after the program exit
// temporary files and directories are useful this purpose since they don't pollute the file system over time

func main() {
	f, err := os.CreateTemp("", "sample")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(f.Name())
	fmt.Println("temp file name: ", f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	if err != nil {
		panic(err)
	}

	dname, err := os.MkdirTemp("", "sampleDir")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dname)
	fmt.Println("temp dir name: ", dname)

	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	if err != nil {
		panic(err)
	}
}
