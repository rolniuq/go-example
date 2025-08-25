package main

import (
	"embed"
	linkname "myembed/link_name"
)

// //go:embed is a compiler directive that allows program include arbitrary files and folders in the Go binary at build time

//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {
	print(fileString)
	print(string(fileByte))

	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))

	sec, nsec := linkname.RuntimeNow()
	print("runtime now", sec, nsec)
}
