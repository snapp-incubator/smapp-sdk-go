package reverse

type ConstructorOption func(client *Client)

func WithURL(url string) ConstructorOption {
	return func(client *Client) {
		client.url = url
	}
}
