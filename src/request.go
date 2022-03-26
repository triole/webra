package main

import (
	"net/http"
	"net/url"
)

func request(targetURL string) (response *http.Response, errs []string) {
	url, err := url.Parse(targetURL)
	if err != nil {
		errs = append(errs, err.Error())
		return
	}

	client := &http.Client{}
	request, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		errs = append(errs, err.Error())
		return
	}

	request.Header.Set("User-Agent", CLI.UserAgent)
	if err != nil {
		errs = append(errs, err.Error())
		return
	}

	response, err = client.Do(request)
	if err != nil {
		errs = append(errs, err.Error())
	}

	return
}
