package unione

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/alexeyco/unione/message"
)

const (
	// Endpoint Unione API endpoint.
	Endpoint = "https://one.unisender.com"

	// DefaultLanguage default API language.
	DefaultLanguage = "en"
)

// Client Unione API client interface.
type Client interface {
	// LanguageEn sets API response language to English (default).
	LanguageEn() Client

	// LanguageRu sets API response language to Russian.
	LanguageRu() Client

	// Client setter for custom http client.
	Client(client *http.Client) Client

	// Send sends transactional email to recipients.
	//
	// See: https://one.unisender.com/en/docs/page/send
	Send(m message.Message) (success []string, failed map[string]error, err error)
}

type client struct {
	userName string
	apiKey   string
	language string
	client   *http.Client
}

func (c *client) LanguageEn() Client {
	return c
}

func (c *client) LanguageRu() Client {
	return c
}

func (c *client) Client(client *http.Client) Client {
	c.client = client
	return c
}

func (c *client) Send(m message.Message) (success []string, failed map[string]error, err error) {
	var res *Response
	if res, err = c.post("transactional/api/v1/email/send.json", "message", m); err != nil {
		return
	}

	success = res.Emails
	failed = mapErrors(res.FailedEmails)

	return
}

func (c *client) post(resource, key string, data interface{}) (response *Response, err error) {
	requestData := map[string]interface{}{
		"username": c.userName,
		"api_key":  c.apiKey,
	}

	requestData[key] = data

	var b []byte
	if b, err = json.Marshal(&requestData); err != nil {
		return
	}

	url := fmt.Sprintf("%s/%s/%s", Endpoint, c.language, resource)

	var req *http.Request
	if req, err = http.NewRequest("POST", url, bytes.NewBuffer(b)); err != nil {
		return
	}

	var res *http.Response
	if res, err = c.client.Do(req); err != nil {
		return
	}
	defer func() {
		_ = res.Body.Close()
	}()

	var body []byte
	if body, err = ioutil.ReadAll(res.Body); err != nil {
		return
	}

	if len(body) == 0 {
		return
	}

	var resp Response
	if err = json.Unmarshal(body, &resp); err != nil {
		return
	}

	response = &resp
	if res.StatusCode != http.StatusOK {
		err = errors.New(response.Message)
	}

	return
}

func mapErrors(src map[string]string) map[string]error {
	errs := map[string]error{}
	for email, e := range src {
		errs[email] = errors.New(e)
	}

	return errs
}

// New returns new Unione API client.
func New(userName, apiKey string) Client {
	return &client{
		userName: userName,
		apiKey:   apiKey,
		language: DefaultLanguage,
		client:   http.DefaultClient,
	}
}
