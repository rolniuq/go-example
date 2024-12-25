package main

import (
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed init.templ
var tmpl []byte

func render(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write(tmpl)
}

func main() {
	http.HandleFunc("/", render)

	http.Handle("/styles.css", http.FileServer(http.Dir(".")))
	http.Handle("/script.js", http.FileServer(http.Dir(".")))

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
