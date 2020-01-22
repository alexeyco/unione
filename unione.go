package unione

import "github.com/alexeyco/unione/message"

type Client struct {
	userName string
	apiKey   string
}

func (c *Client) Send(m message.Message) (err error) {
	return
}

func New(userName, apiKey string) *Client {
	return &Client{
		userName: userName,
		apiKey:   apiKey,
	}
}
