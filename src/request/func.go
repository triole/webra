package request

import (
	"net/http"
	"net/url"
	"time"
)

// HTTP makes a http request
func (req Req) HTTP(targetURL string) (response *http.Response, errs []string) {
	url, err := url.Parse(targetURL)
	if err != nil {
		errs = append(errs, err.Error())
		return
	}

	transport, err := req.setProxy()
	if err != nil {
		errs = append(errs, err.Error())
		return
	}

	client := &http.Client{
		Timeout:   time.Duration(req.Settings.TimeOut) * time.Second,
		Transport: transport,
	}

	request, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		errs = append(errs, err.Error())
		return
	}

	request.Header.Set("User-Agent", req.Settings.UserAgent)
	if err != nil {
		errs = append(errs, err.Error())
		return
	}

	if req.Settings.AuthEnabled == true {
		request.SetBasicAuth(
			req.Settings.AuthUser, req.Settings.AuthPass,
		)
	}

	response, err = client.Do(request)
	if err != nil {
		errs = append(errs, err.Error())
	}

	return
}

func (req *Req) setProxy() (trans http.RoundTripper, err error) {
	var proxyURL *url.URL
	if req.Settings.ProxyURL != "" {
		proxyURL, err = url.Parse(req.Settings.ProxyURL)
		if err != nil {
			return
		}
		trans = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}
	return
}
