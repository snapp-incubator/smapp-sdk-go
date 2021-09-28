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

type CallOptions struct {
	ZoomLevel    int
	ResponseType ResponseType
	Language     Language
}

type CallOptionSetter func(options *CallOptions)

func WithDriverResponseType() CallOptionSetter {
	return func(options *CallOptions) {
		options.ResponseType = Driver
	}
}

func WithPassengerResponseType() CallOptionSetter {
	return func(options *CallOptions) {
		options.ResponseType = Passenger
	}
}

func WithVerboseResponseType() CallOptionSetter {
	return func(options *CallOptions) {
		options.ResponseType = Verbose
	}
}

func WithFarsiLanguage() CallOptionSetter {
	return func(options *CallOptions) {
		options.Language = Farsi
	}
}

func WithEnglishLanguage() CallOptionSetter {
	return func(options *CallOptions) {
		options.Language = English
	}
}

func WithZoomLevel(zoomLevel int) CallOptionSetter {
	return func(options *CallOptions) {
		options.ZoomLevel = zoomLevel
	}
}

func NewDefaultCallOptions(opts ...CallOptionSetter) CallOptions {
	callOptions := CallOptions{
		ZoomLevel:    16,
		ResponseType: Driver,
		Language:     Farsi,
	}

	for _, opt := range opts {
		opt(&callOptions)
	}

	return callOptions
}
