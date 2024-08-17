package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	CopeHandle("GET", "https://google.com", "")
}

func CopeHandle(method string, urlVal string, data string) {
	client := &http.Client{}
	var req *http.Request

	if data == "" {
		urlArr := strings.Split(urlVal, "?")
		if len(urlArr) == 2 {
			urlVal = urlArr[0] + "?" + getParseParam(urlArr[1])
		}
		req, _ = http.NewRequest(method, urlVal, nil)
	} else {
		req, _ = http.NewRequest(method, urlVal, strings.NewReader(data))
	}

	cookie := &http.Cookie{Name: "name", Value: "value", HttpOnly: true}
	req.AddCookie(cookie)

	req.Header.Add("name", "value")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	fmt.Println(string(b))
}

func getParseParam(param string) string {
	return url.PathEscape(param)
}
