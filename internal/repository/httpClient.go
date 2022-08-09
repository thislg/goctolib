package repository

import (
	"fmt"
	"net/http"
)

func GetResponse(url string, method string, verbose bool) *http.Response {
	if verbose {
		fmt.Println(method + " " + url)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	req.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.27 Safari/537.36`)
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	if verbose {
		fmt.Println("Response status:", resp.Status)
	}

	return resp
}
