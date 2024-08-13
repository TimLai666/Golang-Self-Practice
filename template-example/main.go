package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func helloHandleFunc(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./template_example.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	name := "Tim"
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", helloHandleFunc)
	http.ListenAndServe(":8080", nil)
}
