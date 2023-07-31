package param

type Client struct {
	Region string
}

func NewClient() *Client {
	return &Client{}
}
