package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	closer := resp.Body
	bytes, _ := io.ReadAll(closer)
	fmt.Println(string(bytes))
}
