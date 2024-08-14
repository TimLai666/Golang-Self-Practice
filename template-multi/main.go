package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func tmplSample(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./t.html", "./ul.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	user := UserInfo{
		Name:   "張三",
		Gender: "男",
		Age:    18,
	}
	tmpl.Execute(w, user)
}

func main() {
	http.HandleFunc("/", tmplSample)
	http.ListenAndServe(":8080", nil)
}
