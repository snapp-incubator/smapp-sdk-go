package eta

// CallOptions is the type that specifies behaviour of a locate request.
type CallOptions struct {
	// UseNoTraffic specifies if `no_traffic` query param exists in request.
	UseNoTraffic bool
	// NoTraffic is the value of `no_traffic` query param.
	NoTraffic bool
	// UseDepartureDateTime specifies if `departure_date_time` field exists in `json` query parameter.
	UseDepartureDateTime bool
	// DepartureDateTime is the value of `departure_date_time` field in `json` query parameter.
	DepartureDateTime string
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

// WithNoTraffic will set `no_traffic` query param ro true. with this option eta requests does not involve traffic data in response.
func WithNoTraffic() CallOptionSetter {
	return func(options *CallOptions) {
		options.UseNoTraffic = true
		options.NoTraffic = true
	}
}

// WithDepartureDateTime will set the departure date time of the eta request.
func WithDepartureDateTime(dateTime string) CallOptionSetter {
	return func(options *CallOptions) {
		options.UseDepartureDateTime = true
		options.DepartureDateTime = dateTime
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
