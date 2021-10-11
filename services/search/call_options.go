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

// CallOptions is the type that specifies behaviour of a search request.
type CallOptions struct {
	// UseLocation specifies if `location` query param exists in request.
	UseLocation bool
	// Location is the data-structure for keeping a location.
	Location struct {
		Lat float64
		Lon float64
	}
	// UseLanguage specifies if `language` query param exists in request.
	UseLanguage bool
	// Language specifies tha language of response.
	Language Language
	// UseRequestContext specifies if `context` query param exists in request.
	UseRequestContext bool
	// RequestContext specifies context of a search request.
	RequestContext RequestContext
	// UseUserLocation specifies if `user_location` query param exists in request.
	UseUserLocation bool
	// UserLocation is the data-structure for keeping a location.
	UserLocation struct {
		Lat float64
		Lon float64
	}
	// UseCityID specifies if `city_id` query param exists in request.
	UseCityID bool
	// CityID contains city_id of the request
	CityID int
	// Headers is a map that contains all custom headers to be sent.
	Headers map[string]string
}

// CallOptionSetter is a function for defining custom call options in a fluent way.
type CallOptionSetter func(options *CallOptions)

// WithLocation will set a latitude and longitude for a request
func WithLocation(lat, lon float64) CallOptionSetter {
	return func(options *CallOptions) {
		options.UseLocation = true
		options.Location.Lat = lat
		options.Location.Lon = lon
	}
}

// WithFarsiLanguage will set Farsi as response language
func WithFarsiLanguage() CallOptionSetter {
	return func(options *CallOptions) {
		options.UseLanguage = true
		options.Language = Farsi
	}
}

// WithEnglishLanguage will set English as response language
func WithEnglishLanguage() CallOptionSetter {
	return func(options *CallOptions) {
		options.UseLanguage = true
		options.Language = English
	}
}

// WithOriginRequestContext will set Origin as request context.
func WithOriginRequestContext() CallOptionSetter {
	return func(options *CallOptions) {
		options.UseRequestContext = true
		options.RequestContext = Origin
	}
}

// WithFavouriteRequestContext will set WithFavouriteRequestContext as request context.
func WithFavouriteRequestContext() CallOptionSetter {
	return func(options *CallOptions) {
		options.UseRequestContext = true
		options.RequestContext = Favourite
	}
}

// WithFirstDestinationRequestContext will set FirstDestination as request context.
func WithFirstDestinationRequestContext() CallOptionSetter {
	return func(options *CallOptions) {
		options.UseRequestContext = true
		options.RequestContext = FirstDestination
	}
}

// WithSecondDestinationRequestContext will set SecondDestination as request context.
func WithSecondDestinationRequestContext() CallOptionSetter {
	return func(options *CallOptions) {
		options.UseRequestContext = true
		options.RequestContext = SecondDestination
	}
}

// WithUserLocation will set a latitude and longitude for a request as user's location
func WithUserLocation(lat, lon float64) CallOptionSetter {
	return func(options *CallOptions) {
		options.UseUserLocation = true
		options.UserLocation.Lat = lat
		options.UserLocation.Lon = lon
	}
}

// WithCityId will set the given id for `city_id` query param in request.
func WithCityId(cityId int) CallOptionSetter {
	return func(options *CallOptions) {
		options.UseCityID = true
		options.CityID = cityId
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
		UseLocation:       false,
		UseLanguage:       false,
		UseRequestContext: false,
		Language:          Farsi,
		RequestContext:    Origin,
		UseUserLocation:   false,
		UseCityID:         false,
		Headers:           make(map[string]string),
	}

	for _, opt := range opts {
		opt(&callOptions)
	}

	return callOptions
}
