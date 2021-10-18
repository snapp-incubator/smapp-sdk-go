package reverse

import "net/http"

// ConstructorOption is a function type for customizing constructor behaviour in a fluent way.
type ConstructorOption func(client *Client)

// WithURL will override base url of the service.
func WithURL(url string) ConstructorOption {
	return func(client *Client) {
		client.url = url
	}
}

// WithTransport will override request http client transport object
func WithTransport(transport *http.Transport) ConstructorOption {
	return func(client *Client) {
		client.httpClient.Transport = transport
	}
}
