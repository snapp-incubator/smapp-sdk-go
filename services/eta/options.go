package eta

import (
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"net/http"
)

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

// WithRequestOpenTelemetryTracing will enable opentelemetry tracing the request
func WithRequestOpenTelemetryTracing(tracerName string) ConstructorOption {
	return func(client *Client) {
		client.tracerName = tracerName
		client.httpClient.Transport = otelhttp.NewTransport(client.httpClient.Transport)
	}
}