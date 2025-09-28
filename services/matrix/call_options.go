package matrix

// MatrixEngine type is for defining different engines
// that can be used in calculating the matrix eta.
type MatrixEngine int

const (
	MatrixEngineV1 MatrixEngine = iota
	MatrixEngineV2

	MatrixEngineOcelot
	MatrixEngineOrca
	MatrixEngineOrcaCh

	MatrixEngineMurche
	MatrixEngineKandoo
	MatrixEngineZeus
	MatrixEngineCarPooling
	MatrixEngineGolchin
	MatrixEngineGiraffe
	MatrixEngineFood
	MatrixEngineIntercity
)

// String casts the engine enum to its string value.
func (engine MatrixEngine) String() string {
	switch engine {
	case MatrixEngineV1:
		return "v1"
	case MatrixEngineV2:
		return "v2"

	case MatrixEngineOcelot:
		return "ocelot"
	case MatrixEngineOrca:
		return "orca"
	case MatrixEngineOrcaCh:
		return "orca-ch"

	case MatrixEngineMurche:
		return "murche"
	case MatrixEngineKandoo:
		return "kandoo"
	case MatrixEngineZeus:
		return "zeus"
	case MatrixEngineCarPooling:
		return "carpooling"
	case MatrixEngineGolchin:
		return "golchin"
	case MatrixEngineGiraffe:
		return "giraffe"
	case MatrixEngineFood:
		return "food"
	case MatrixEngineIntercity:
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
	// Engine is the value of `engine` query param.
	Engine MatrixEngine
	// Engine is the value of `engine` in query param as string.
	EngineStr string
	// Headers is a map that contains all custom headers to be sent.
	Headers map[string]string
	// UsePost to use post http call for bigger matrix calls
	UsePost bool
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

// WithUsePost will use post http call for bigger matrix calls
func WithUsePost() CallOptionSetter {
	return func(options *CallOptions) {
		options.UsePost = true
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

// WithEngine will set `engine` query param to passed value.
func WithEngine(engine MatrixEngine) CallOptionSetter {
	return func(options *CallOptions) {
		options.Engine = engine
	}
}

func WithEngineStr(engineStr string) CallOptionSetter {
	return func(options *CallOptions) {
		options.EngineStr = engineStr
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
