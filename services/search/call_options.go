package search

type Language string
type RequestContext string

const (
	English Language = "en"
	Farsi   Language = "fa"

	Origin            RequestContext = "origin"
	Favourite         RequestContext = "favourite"
	FirstDestination  RequestContext = "destination1"
	SecondDestination RequestContext = "destination2"
)

type CallOptions struct {
	UseLocation bool
	Location    struct {
		Lat float64
		Lon float64
	}
	Language        Language
	RequestContext  RequestContext
	UseUserLocation bool
	UserLocation    struct {
		Lat float64
		Lon float64
	}
	UseCityID bool
	CityID    int
	Headers   map[string]string
}

type CallOptionSetter func(options *CallOptions)

func WithLocation(lat, lon float64) CallOptionSetter {
	return func(options *CallOptions) {
		options.UseLocation = true
		options.Location.Lat = lat
		options.Location.Lon = lon
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

func WithOriginRequestContext() CallOptionSetter {
	return func(options *CallOptions) {
		options.RequestContext = Origin
	}
}

func WithFavouriteRequestContext() CallOptionSetter {
	return func(options *CallOptions) {
		options.RequestContext = Favourite
	}
}

func WithFirstDestinationRequestContext() CallOptionSetter {
	return func(options *CallOptions) {
		options.RequestContext = FirstDestination
	}
}

func WithSecondDestinationRequestContext() CallOptionSetter {
	return func(options *CallOptions) {
		options.RequestContext = SecondDestination
	}
}

func WithUserLocation(lat, lon float64) CallOptionSetter {
	return func(options *CallOptions) {
		options.UseUserLocation = true
		options.UserLocation.Lat = lat
		options.UserLocation.Lon = lon
	}
}

func WithCityId(cityId int) CallOptionSetter {
	return func(options *CallOptions) {
		options.UseCityID = true
		options.CityID = cityId
	}
}

func WithHeaders(headers map[string]string) CallOptionSetter {
	return func(options *CallOptions) {
		if headers != nil {
			options.Headers = headers
		}
	}
}

func NewDefaultCallOptions(opts ...CallOptionSetter) CallOptions {
	callOptions := CallOptions{
		UseLocation:     false,
		Language:        Farsi,
		RequestContext:  Origin,
		UseUserLocation: false,
		UseCityID: false,
		Headers:   make(map[string]string),
	}

	for _, opt := range opts {
		opt(&callOptions)
	}

	return callOptions
}
