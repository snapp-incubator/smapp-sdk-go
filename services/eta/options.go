package eta

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// PathStyle determines how the service path/version is combined with the base URL.
// LegacyPathStyle is the legacy layout: {base}/eta/{version}
// V1PathStyle is the new layout:     {base}/api/{version}/eta
type PathStyle int

const (
	LegacyPathStyle PathStyle = iota
	V1PathStyle
)

// ConstructorOption is a function type for customizing constructor behaviour in a fluent way.
type ConstructorOption func(client *Client)

// WithURL will override base url of the service.
func WithURL(url string) ConstructorOption {
	return func(client *Client) {
		client.url = url
	}
}

// WithPathStyle will override the default path style used to build the service URL.
// MonshiPathStyle: {base}/eta/{version}
// BifrostPathStyle: {base}/api/{version}/eta
func WithPathStyle(style PathStyle) ConstructorOption {
	return func(client *Client) {
		client.pathStyle = style
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
