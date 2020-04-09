package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", HelloworldHandler)
	http.ListenAndServe(":80", nil)
}

func HelloworldHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World!")
}