package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request) {
	panic("error")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, v interface{}) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "An error occurred: %s", v)
	}
	log.Fatal(http.ListenAndServe(":8080", router))
}
