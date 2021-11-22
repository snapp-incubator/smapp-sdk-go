package area_gateways

type Language string

const (
	Farsi   Language = "fa"
	English Language = "en"
)

// CallOptions is the type that specifies behaviour of a area-gateways request.
type CallOptions struct {
	// UseLanguage specifies if `Accept-Language` header exists in request.
	UseLanguage bool
	// Language of the response
	Language Language
	// Headers is a map that contains all custom headers to be sent.
	Headers map[string]string
}

// CallOptionSetter is a function for defining custom call options in a fluent way.
type CallOptionSetter func(options *CallOptions)

// WithFarsiLanguage will set the response language to Farsi
func WithFarsiLanguage() CallOptionSetter {
	return func(options *CallOptions) {
		options.UseLanguage = true
		options.Language = Farsi
	}
}

// WithEnglishLanguage will set the response language to English
func WithEnglishLanguage() CallOptionSetter {
	return func(options *CallOptions) {
		options.UseLanguage = true
		options.Language = English
	}
}

// WithHeaders will set given header map to extra headers to be sent in request
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
		UseLanguage: false,
		Language:    Farsi,
		Headers:     make(map[string]string),
	}

	for _, opt := range opts {
		opt(&callOptions)
	}

	return callOptions
}
