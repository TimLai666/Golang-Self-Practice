package main

import (
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{Addr: ":8080", Handler: http.HandlerFunc(handle)}
	log.Printf("Server started on port 8080")
	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
}

func handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request received from %s", r.Proto)
	w.Write([]byte("Hello, World!"))
}
