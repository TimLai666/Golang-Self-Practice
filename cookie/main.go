package main

import (
	"fmt"
	"net/http"
)

func testHandle(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("test")
	fmt.Printf("cookie: %v, err: %v\n", c, err)

	cookie := &http.Cookie{
		Name:   "test",
		Value:  "test",
		MaxAge: 3600,
		Domain: "localhost",
		Path:   "/",
	}

	http.SetCookie(w, cookie)

	w.Write([]byte("test"))
}

func main() {
	http.HandleFunc("/test", testHandle)
	http.ListenAndServe(":8080", nil)
}
