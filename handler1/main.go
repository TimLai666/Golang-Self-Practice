package main

import (
	"fmt"
	"net/http"
)

type handle1 struct{}

func (h1 *handle1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "handle1")
}

type handle2 struct{}

func (h2 *handle2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "handle2")
}

func main() {
	handle1 := handle1{}
	handle2 := handle2{}

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: nil,
	}

	http.Handle("/handle1", &handle1)
	http.Handle("/handle2", &handle2)
	server.ListenAndServe()
}
