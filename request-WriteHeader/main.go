package main

import (
	"fmt"
	"net/http"
)

func noAuth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(401)
	fmt.Fprintln(w, "401 Unauthorized")
}

func main() {
	http.HandleFunc("/no-auth", noAuth)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
