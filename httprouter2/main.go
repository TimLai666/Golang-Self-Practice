package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HostMap map[string]http.Handler

func (h HostMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler := h[r.Host]; handler != nil {
		handler.ServeHTTP(w, r)
	} else {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}
}

func main() {
	userRouter := httprouter.New()
	userRouter.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("Welcome User"))
	})

	adminRouter := httprouter.New()
	adminRouter.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("Welcome Admin"))
	})

	hs := make(HostMap)

	hs["user.example.com"] = userRouter
	hs["admin.example.com"] = adminRouter

	log.Fatal(http.ListenAndServe(":8080", hs))
}
