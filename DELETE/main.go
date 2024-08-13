package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	url := "https://www.shirdon.com/comment/delete"
	payload := strings.NewReader("{\"userId\":1, \"articleId\":1, \"comment\":\"This is a comment\"}")
	req, _ := http.NewRequest("DELETE", url, payload)
	req.Header.Add("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
