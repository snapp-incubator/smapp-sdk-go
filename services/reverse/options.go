package reverse

// ConstructorOption is a function type for customizing constructor behaviour in a fluent way.
type ConstructorOption func(client *Client)

// WithURL will override base url of the service.
func WithURL(url string) ConstructorOption {
	return func(client *Client) {
		client.url = url
	}
}
