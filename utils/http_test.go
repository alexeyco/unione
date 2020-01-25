package utils_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/alexeyco/unione/utils"
)

func TestNewTestHttpClient(t *testing.T) {
	expectedRequestMethod := "POST"
	var givenRequestMethod string

	expectedRequestUrl := "https://foo.bar/baz"
	var givenRequestUrl string

	expectedRequestJson := `{"foo":{"bar":"baz"}}`
	var givenRequestJson string

	expectedResponseStatus := http.StatusTeapot
	var givenResponseStatus int

	expectedResponseJson := `{"message":"I'm batman!'"}`
	var givenResponseJson string

	client := utils.NewTestHttpClient(func(req *http.Request) (res *http.Response, err error) {
		givenRequestMethod = req.Method
		givenRequestUrl = req.URL.String()

		var b []byte
		if b, err = ioutil.ReadAll(req.Body); err != nil {
			return
		}

		givenRequestJson = string(b)

		res = &http.Response{
			StatusCode: expectedResponseStatus,
			Body:       ioutil.NopCloser(bytes.NewBufferString(expectedResponseJson)),
		}

		return
	})

	req, err := http.NewRequest(expectedRequestMethod, expectedRequestUrl, ioutil.NopCloser(bytes.NewBufferString(expectedRequestJson)))
	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err)
	}

	var res *http.Response
	if res, err = client.Do(req); err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err)
	}

	givenResponseStatus = res.StatusCode

	var b []byte
	if b, err = ioutil.ReadAll(res.Body); err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err)
	}

	givenResponseJson = string(b)

	if expectedRequestMethod != givenRequestMethod {
		t.Fatalf(`Request method should be "%s", "%s" given`, expectedRequestMethod, givenRequestMethod)
	}

	if expectedRequestUrl != givenRequestUrl {
		t.Fatalf(`Request URL should be "%s", "%s" given`, expectedRequestUrl, givenRequestUrl)
	}

	if expectedRequestJson != givenRequestJson {
		t.Fatalf(`Request JSON should be %s, %s given`, expectedRequestJson, givenRequestJson)
	}

	if expectedResponseStatus != givenResponseStatus {
		t.Fatalf(`Response status should be %d, %d given`, expectedResponseStatus, givenResponseStatus)
	}

	if expectedResponseJson != givenResponseJson {
		t.Fatalf(`Response JSON should be %s, %s given`, expectedResponseJson, givenResponseJson)
	}
}
