package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/default", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("default GET"))
	})

	router.POST("/default", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("default POST"))
	})

	// router.GET("/user/:name", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 	w.Write([]byte("user " + ps.ByName("name")))
	// })

	router.GET("/user/*name", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Write([]byte("user " + ps.ByName("name")))
	})

	http.ListenAndServe(":8080", router)
}
