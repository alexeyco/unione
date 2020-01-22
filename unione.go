package unione

type Client struct {
	userName string
	apiKey   string
}

func New(userName, apiKey string) *Client {
	return &Client{
		userName: userName,
		apiKey:   apiKey,
	}
}
