package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, World!</h1>")
}

func main() {
	server := &http.Server{
		Addr: ":80",
	}
	http.HandleFunc("/", hello)
	server.ListenAndServe()
}
