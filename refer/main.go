package main

import "net/http"

type Refer struct {
	handler http.Handler
	refer   string
}

func (this *Refer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Referer() != this.refer {
		w.WriteHeader(403)
		return
	}
	this.handler.ServeHTTP(w, r)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a handler"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	referer := &Refer{
		handler: http.HandlerFunc(myHandler),
		refer:   "http://localhost:8080",
	}
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", referer)
}
