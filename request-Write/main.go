package main

import "net/http"

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}

func main() {
	http.HandleFunc("/", Welcome)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
