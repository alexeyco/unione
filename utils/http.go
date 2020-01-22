package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type RoundTripHandler func(*http.Request) (*http.Response, error)

type roundTripper struct {
	handler RoundTripHandler
}

func (r *roundTripper) RoundTrip(req *http.Request) (res *http.Response, err error) {
	res, err = r.handler(req)
	if res == nil {
		res = &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(""))),
		}
	}

	return
}

func NewHttpClient(h RoundTripHandler) (client *http.Client) {
	client = http.DefaultClient
	client.Transport = &roundTripper{
		handler: h,
	}

	return
}
