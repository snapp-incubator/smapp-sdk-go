package eta

// EtaEngine type is for defining different engines
// that can be used in calculating the eta.
type EtaEngine int

const (
	EtaEngineV1 EtaEngine = iota
	EtaEngineV2

	EtaEngineNostradamus
	EtaEngineOcelot
	EtaEngineOrca

	EtaEngineMurche
	EtaEngineKandoo
	EtaEngineZeus
	EtaEngineCarPooling
	EtaEngineGolchin
	EtaEngineGiraffe
	EtaEngineFood
	EtaEngineIntercity
)

// String casts the engine enum to its string value.
func (engine EtaEngine) String() string {
	switch engine {
	case EtaEngineV1:
		return "v1"
	case EtaEngineV2:
		return "v2"

	case EtaEngineNostradamus:
		return "nostradamus"
	case EtaEngineOcelot:
		return "ocelot"
	case EtaEngineOrca:
		return "orca"

	case EtaEngineMurche:
		return "murche"
	case EtaEngineKandoo:
		return "kandoo"
	case EtaEngineZeus:
		return "zeus"
	case EtaEngineCarPooling:
		return "carpooling"
	case EtaEngineGolchin:
		return "golchin"
	case EtaEngineGiraffe:
		return "giraffe"
	case EtaEngineFood:
		return "food"
	case EtaEngineIntercity:
		return "intercity"
	}
	return "v1"
}

// CallOptions is the type that specifies behaviour of a eta request.
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
	// Engine is the value of `engine` query param.
	Engine EtaEngine
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

// WithTraffic will set `no_traffic` query param ro false. with this option eta requests does involve traffic data in response.
func WithTraffic() CallOptionSetter {
	return func(options *CallOptions) {
		options.UseNoTraffic = true
		options.NoTraffic = false
	}
}

// WithDepartureDateTime will set the departure date time of the eta request.
func WithDepartureDateTime(dateTime string) CallOptionSetter {
	return func(options *CallOptions) {
		options.UseDepartureDateTime = true
		options.DepartureDateTime = dateTime
	}
}

// WithEngine will set `engine` query param to passed value.
func WithEngine(engine EtaEngine) CallOptionSetter {
	return func(options *CallOptions) {
		options.Engine = engine
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
