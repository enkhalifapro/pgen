package utilities

import (
	"bytes"
	"net/http"
)

type HttpClientUtil struct {
}

func (h *HttpClientUtil) Post(url string, body []byte) (response *http.Response, err error) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("charset", "utf-8")
	hc := http.Client{}
	return hc.Do(req)
}
