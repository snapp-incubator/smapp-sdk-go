package matrix


// CallOptions is the type that specifies behaviour of a eta request.
type CallOptions struct {
	// Headers is a map that contains all custom headers to be sent.
	Headers map[string]string
}

// CallOptionSetter is a function for defining custom call options in a fluent way.
type CallOptionSetter func(options *CallOptions)

// WithHeaders will set given header map to extra headers to be sent in request.
func WithHeaders(headers map[string]string) CallOptionSetter {
	return func(options *CallOptions) {
		if headers != nil {
			options.Headers = headers
		}
	}
}

// NewDefaultCallOptions is the constructor of a default CallOptions
func NewDefaultCallOptions(opts ...CallOptionSetter) CallOptions {
	callOptions := CallOptions{
		Headers: make(map[string]string),
	}

	for _, opt := range opts {
		opt(&callOptions)
	}

	return callOptions
}

