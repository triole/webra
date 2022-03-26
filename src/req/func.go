package req

import (
	"net/http"
	"net/url"
)

// HTTP makes a http request
func (req Req) HTTP(targetURL string) (response *http.Response, errs []string) {
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

	request.Header.Set("User-Agent", req.HTTPUserAgent)
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
