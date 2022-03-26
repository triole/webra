package main

import (
	"net/http"
	"net/url"
)

func request(targetURL string) (response *http.Response) {
	url, err := url.Parse(targetURL)
	printErr(err)
	client := &http.Client{}

	request, err := http.NewRequest("GET", url.String(), nil)
	request.Header.Set("User-Agent", CLI.UserAgent)
	printErr(err)

	response, err = client.Do(request)
	printErr(err)

	return
}
