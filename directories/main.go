package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	// create subdir folder
	err := os.Mkdir("subdir", 0755)
	if err != nil {
		panic(err)
	}
	// remove subdir folder before func exit
	defer os.RemoveAll("subdir")

	createEmptyFile := func(name string) {
		d := []byte("")
		err := os.WriteFile(name, d, 0644)
		if err != nil {
			panic(err)
		}
	}
	// create empty file1
	createEmptyFile("subdir/file1")

	// create folder child
	err = os.MkdirAll("subdir/parent/child", 0755)
	if err != nil {
		panic(err)
	}

	// create empty file 2
	createEmptyFile("subdir/parent/file2")
	// create empty file 3
	createEmptyFile("subdiflderr/parent/file3")
	// create empty file 4
	createEmptyFile("subdir/parent/child/file4")

	// read dir
	c, err := os.ReadDir("subdir/parent")
	if err != nil {
		panic(err)
	}

	fmt.Println("listen subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// provide role
	err = os.Chdir("subdir/parent/child")
	if err != nil {
		panic(err)
	}

	c, err = os.ReadDir(".")
	if err != nil {
		panic(err)
	}

	fmt.Println("listent subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	if err != nil {
		panic(err)
	}

	fmt.Println("visit subdir")
	err = filepath.WalkDir("subdir", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(" ", path, d.IsDir())
		return nil
	})
	if err != nil {
		panic(err)
	}
}
