package main

import (
	"fmt"
	"net/http"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://www.google.com")
	w.WriteHeader(301)
}

func main() {
	http.HandleFunc("/", Redirect)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
