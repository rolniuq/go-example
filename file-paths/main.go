package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// the filepath pkg provide functions to parse and construct file paths in a way that is portable between operating system
//  dir/file on linux vs dir/file on windows

func main() {
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p: ", p)

	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	fmt.Println("dir(p): ", filepath.Dir(p))
	fmt.Println("base(p): ", filepath.Base(p))

	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println("ext: ", ext)

	fmt.Println("trum suffix: ", strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println("rel: ", rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println("rel: ", rel)
}
