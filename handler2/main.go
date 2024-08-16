package main

import (
	"fmt"
	"log"
	"net/http"
)

type WelcomeHandler struct {
	Language string
}

func (h WelcomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", h.Language)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/zh", WelcomeHandler{Language: "歡迎"})
	mux.Handle("/en", WelcomeHandler{Language: "Welcome"})

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
