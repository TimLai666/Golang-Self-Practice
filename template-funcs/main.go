package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func Welcome() string {
	return "Welcome to Go Web Development"
}

func Doing(name string) string {
	return name + " is learning how to use templates"
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	htmlByte, err := os.ReadFile("./funcs.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	loveGo := func() string {
		return "I love Go!"
	}
	tmpl1, err := template.New("funcs").Funcs(template.FuncMap{"loveGo": loveGo}).Parse(string(htmlByte))
	if err != nil {
		fmt.Println(err)
		return
	}
	funcMap := template.FuncMap{
		"Welcome": Welcome,
		"Doing":   Doing,
	}
	name := "Tim"
	tmpl2, err := template.New("test").Funcs(funcMap).Parse("{{Welcome}}\n{{Doing .}}\n")
	if err != nil {
		panic(err)
	}

	tmpl1.Execute(w, name)
	tmpl2.Execute(w, name)
}

func main() {
	http.HandleFunc("/", sayHello)
	http.ListenAndServe(":8080", nil)
}
