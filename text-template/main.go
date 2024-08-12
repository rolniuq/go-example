package main

import (
	"os"
	"text/template"
)

func main() {
	t1 := template.Must(template.New("test").Parse("hello {{.}}"))
	t1.Execute(os.Stdout, "world\n")

	t2 := template.Must(template.New("condition").Parse(`{{if .}}true{{else}}false{{end}}`))
	t2.Execute(os.Stdout, "valid\n")
}
