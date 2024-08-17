package main

import (
	"encoding/json"
	"net/http"
)

type Greeting struct {
	Msg string `json:"msg"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	greeting := Greeting{Msg: "Hello, World!"}
	message, _ := json.Marshal(greeting)
	w.Header().Set("Content-Type", "application/json")
	w.Write(message)
}

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}
