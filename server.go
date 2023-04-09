package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	url := "https://reqres.in/api/users"

	payload := strings.NewReader("name=test&jab=teacher&gender=mi24")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}
