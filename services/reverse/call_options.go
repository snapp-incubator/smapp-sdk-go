package reverse

type ResponseType string
type Language string

const (
	Driver    ResponseType = "driver"
	Passenger ResponseType = "passenger"
	Verbose   ResponseType = "verbose"
)

const (
	Farsi   Language = "fa"
	English Language = "en"
)

// CallOptions is the type that specifies behaviour of a reverse geocode request.
type CallOptions struct {
	UseZoomLevel    bool
	ZoomLevel       int
	UseResponseType bool
	ResponseType    ResponseType
	UseLanguage     bool
	Language        Language
}

// CallOptionSetter is a function for defining custom call options in a fluent way.
type CallOptionSetter func(options *CallOptions)

// WithDriverResponseType will set `driver` type for the response
func WithDriverResponseType() CallOptionSetter {
	return func(options *CallOptions) {
		options.UseResponseType = true
		options.ResponseType = Driver
	}
}

// WithPassengerResponseType will set `passenger` type for the response
func WithPassengerResponseType() CallOptionSetter {
	return func(options *CallOptions) {
		options.UseResponseType = true
		options.ResponseType = Passenger
	}
}

// WithVerboseResponseType will set `verbose` type for the response
func WithVerboseResponseType() CallOptionSetter {
	return func(options *CallOptions) {
		options.UseResponseType = true
		options.ResponseType = Verbose
	}
}

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

// WithZoomLevel will set the given zoom level for the request.
func WithZoomLevel(zoomLevel int) CallOptionSetter {
	return func(options *CallOptions) {
		options.UseZoomLevel = true
		options.ZoomLevel = zoomLevel
	}
}

// NewDefaultCallOptions is the constructor of a default CallOptions
func NewDefaultCallOptions(opts ...CallOptionSetter) CallOptions {
	callOptions := CallOptions{
		UseZoomLevel:    false,
		UseResponseType: false,
		UseLanguage:     false,
		ZoomLevel:       16,
		ResponseType:    Driver,
		Language:        Farsi,
	}

	for _, opt := range opts {
		opt(&callOptions)
	}

	return callOptions
}
