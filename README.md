# Smapp Golang SDK

[![pipeline status](https://gitlab.snapp.ir/Map/sdk/smapp-sdk-go/badges/master/pipeline.svg)](https://gitlab.snapp.ir/Map/sdk/smapp-sdk-go/-/commits/master)
[![coverage report](https://gitlab.snapp.ir/Map/sdk/smapp-sdk-go/badges/master/coverage.svg)](https://gitlab.snapp.ir/Map/sdk/smapp-sdk-go/-/commits/master)

[[_TOC_]]

# Overview

`smapp-sdk-go` is a golang sdk for using different Smapp services.

List of supported services are:

- [x] Reverse Geocode
- [x] Search
- [x] Area Gateways
- [x] ETA
- [x] Matrix
- [ ] Routing

# How to Install

First set up your device to be able to use private golang packages from `gitlab.snapp.ir`. for more information, check
this [wiki](https://gitlab.snapp.ir/Map/sdk/smapp-sdk-go/-/wikis/Private-Golang-Modules)

Then add this dependency in the format below to `go.mod` file.

```
module YOUR_MODULE

go 1.21

require (
    ...
	gitlab.snapp.ir/Map/sdk/smapp-sdk-go v0.9.6
    ...
)

replace gitlab.snapp.ir/Map/sdk/smapp-sdk-go => gitlab.snapp.ir/Map/sdk/smapp-sdk-go.git v0.9.6
```

you can download the library with `go mod download` command.

> **Note**: Versions before `v0.9.5` will be deprecated soon.

# Usage

## Config

for using Smapp SDK, you should instantiate a config first. a config object consists of common properties needed for
using in different services.

configs can be instantiated with code or with environment variables.

import package `gitlab.snapp.ir/Map/sdk/smapp-sdk-go/config` for using it.

### Environment variables

List of environment variables are:

+ **`SMAPP_API_KEY`**: specifies api key for using smapp services. this environment variable is **`required`**.
+ `SMAPP_API_KEY_SOURCE`: specifies the source of api key in a request. valid values are `header` and `query`. default
  value is `header`.
+ `SMAPP_API_KEY_NAME`: specifies the name of api key parameter (header name of query param name) in a request. default value is `X-Smapp-Key`.
+ `SMAPP_API_REGION`: the region of smapp api should be specified here. valid values are `teh-1` and `teh-2`. default
  value is `teh-1`
+ `SMAPP_API_BASE_URL`: is the base url of smapp services (API Gateway). default value
  is `http://smapp-api.apps.inter-dc.okd4.teh-1.snappcloud.io`

a config from environment could be instantiated like code below:

```go
config, err := config.ReadFromEnvironment()
if err != nil {
  panic(err)
}
```

### Code

these are default values of each config field:

+ Region: `teh-1`
+ APIKeySource: `header`
+ APIKeyName: `X-Smapp-Key`
+ APIBaseURL: `http://smapp-api.apps.inter-dc.okd4.teh-1.snappcloud.io`

a config could be instantiated like code below:

```go
config, err := config.NewDefaultConfig("api-key")
if err != nil {
  panic(err)
}
```

### Overriding config

You may want to override some fields of the config object. you can pass multiple options to constructors for overriding
config object.

List of options are:

+ [WithRegion(region string)](#) : sets a region for the config.
+ [WithAPIKey(apikey string)](#) : sets the APIKey for the config. it is often used as an option
  in `ReadFromEnvironment`.
+ [WithAPIBaseURL(baseURL string)](#) : sets a custom base URL for API Gateway of smapp services.
+ [WithAPIKeySource(source APIKeySource)](#) : sets an APIKeySource for the config.
+ [WithAPIKeyName(name string)](#) : sets an APIKeyName for the config
+ [WithPublicURL()](#) : sets the APIBaseURL to public routes of smapp. Notice: make sure you set region before using
  this option. if not set `teh-1` region would be used as default region.
+ [WithInternalURL()](#) : sets the APIBaseURL to internal routes of smapp. Notice: make sure you set region before
  using this option. if not set `teh-1` region would be used as default region.

Example:

```go
cfg, err := config.ReadFromEnvironment(
  config.WithRegion("teh-2"),
  config.WithAPIKey("example-api-key"),
  config.WithPublicURL(),
)

cfg2, err := NewDefaultConfig("api-key",
  config.WithRegion("teh-2"),
  config.WithAPIKeyName("X-Smapp-New-Key"),
  config.WithPublicURL(),
)
```

## Reverse Geocode

After creating a config object, you can construct a reverse geo-code client for your services.

The constructor of Reverse Geocode receives a config, version, timeout and multiple optional options.

some constructor options are:
+ [WithURL(url string)](#): could be used to override url of the service.
+ [WithTransport(transport *http.Transport)](#): could be used to set a new `http.Transport` for http client.
+ [WithRequestOpenTelemetryTracing(tracerName string)](#): could be used for activating opentelemetry tracing on client. for more info see [Here](#opentelemetry-tracing)

Example:

```go
package main

import (
	"fmt"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/config"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/services/reverse"
	"time"
)

func main() {
	cfg, err := config.NewDefaultConfig("api-key")
	if err != nil {
		panic(err)
	}

	reverseClient, err := reverse.NewReverseClient(cfg, reverse.V1, time.Second,
		reverse.WithURL("https://new-url.com"), // This is optional
	)
	if err != nil {
		panic(err)
	}

	displayName, err := reverseClient.GetDisplayName(35.0123, 53.12312, reverse.NewDefaultCallOptions())
	if err != nil {
		panic(err)
	}

	fmt.Println(displayName)
}
```

### Operations

List of operations on a reverse geocode client are:

+ **`GetComponents(lat, lon float64, options CallOptions) ([]Component, error)`**:
  receives `lat`,`lon` and `CallOptions` and returns components of address of location given.
+ **`GetComponentsWithContext(ctx context.Context, lat, lon float64, options CallOptions) ([]Component, error)`**:
  same as `GetComponents` but you can pass your own context for more control if needed.
+ **`GetDisplayName(lat, lon float64, options CallOptions) (string, error)`**:
  receives `lat`,`lon` and `CallOptions` and returns a string as address of given location.
+ **`GetDisplayNameWithContext(ctx context.Context, lat, lon float64, options CallOptions) (string, error)`**:
  same as `GetDisplayName` but you can pass your own context for more control if needed.
+ **`GetFrequent(lat, lon float64, options CallOptions) (FrequentAddress, error)`**:
  receives `lat`, `lon` as a location and CallOptions and returns FrequentAddress for the given location.
+ **`GetFrequentWithContext(ctx context.Context, lat, lon float64, options CallOptions) (FrequentAddress, error)`**:
  same as `GetFrequent` but you can pass your own context for more control if needed.
+ **`GetBatch(request BatchReverseRequest) ([]Result, error)`**
  receives a slice of  Request s and returns Component s of address of location given.
  Does not support type 'frequent' in requests and Does not support type Display option as True
+ **`GetBatchWithContext(ctx context.Context, request BatchReverseRequest) ([]Result, error)`**
  same as `GetBatch` but you can pass your own context for more control if needed.
+ **`GetBatchDisplayName(request BatchReverseRequest) ([]Result, error)`**
  receives a slice of  Request s and returns Component s of address of location given, with only the DisplayName
  Only works when Display is true
+ **`GetBatchDisplayNameWithContext(ctx context.Context, request BatchReverseRequest) ([]Result, error)`**
  same as `GetBatchDisplayName` but you can pass your own context for more control if needed.


### CallOptions

`CallOptions` is a struct that defines the behaviour of the operation. you can create a new `CallOptions`
with `reverse.NewDefaultCallOptions()`
function. you can customize the behaviour with passing multiple call options to the constructor.

list of call options for reverse-geocode service are:

+ [WithZoomLevel(zoomLevel int)](#): sets the zoom level for the request. default value is `16`
+ [WithEnglishLanguage()](#): sets the language for response to English. default is `fa` (Farsi)
+ [WithFarsiLanguage()](#): sets the language for response to Farsi.
+ [WithArabicLanguage()](#): sets the language for response to Arabic. 
+ [WithPassengerResponseType()](#): sets the response type suitable for passengers.
+ [WithDriverResponseType()](#): sets the response type suitable for drivers.
+ [WithVerboseResponseType()](#): sets the response type to a verbose response.
+ [WithBikerResponseType()](#): sets the response type suitable for bikers
+ [WithOriginResponseType()](#): sets the response type for origin points.
+ [WithDestinationResponseType()](#): sets the response type for destination points.
+ [WithIraqResponseType()](#): sets the response type suitable for iraq country. it should be used just for **Baly** product.
+ [WithHeaders(headers map[string]string)](#): sets custom headers for request.

This example will set zoom level of 17 with English response:

```go
displayName, err := reverseClient.GetDisplayName(35.0123, 53.12312, reverse.NewDefaultCallOptions(
  reverse.WithZoomLevel(17),
  reverse.WithEnglishLanguage(),
))

if err != nil {
  panic(err)
}
```

## Search

After creating a config object, you can construct a search client for your services.

The constructor of Search receives a config, version, timeout and multiple optional options.

some constructor options are:
+ [WithURL(url string)](#): could be used to override url of the service.
+ [WithTransport(transport *http.Transport)](#): could be used to set a new `http.Transport` for http client.
+ [WithRequestOpenTelemetryTracing(tracerName string)](#): could be used for activating opentelemetry tracing on client. for more info see [Here](#opentelemetry-tracing)

Example:

```go
package main

import (
	"fmt"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/config"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/services/search"
	"time"
)

func main() {
	cfg, err := config.NewDefaultConfig("api-key")
	if err != nil {
		panic(err)
	}

	searchClient, err := search.NewSearchClient(cfg, search.V1, time.Second,
		search.WithURL("https://new-url.com"), // This is optional
	)
	if err != nil {
		panic(err)
	}

	results, err := searchClient.AutoComplete("example", search.NewDefaultCallOptions(
		search.WithCityId(1000),
		search.WithLocation(35.012, 53.1253),
	))
	if err != nil {
		panic(err)
	}

	fmt.Println(results)
}

```

### Operations

List of operations on a search client are:

+ **`GetCities(options CallOptions) ([]City, error)`**:
  it receives call options and returns list of popular cities.
+ **`GetCitiesWithContext(ctx context.Context, options CallOptions) ([]City, error)`**:
  same as `GetCities` but you can pass your context for more control.
+ **`SearchCity(input string, options CallOptions) ([]City, error)`**:
  receives an input string for search and callOptions and returns list of cities according to input string.
+ **`SearchCityWithContext(ctx context.Context, input string, options CallOptions) ([]City, error)`**:
  same as `SearchCity` but you can pass your context for more control.
+ **`AutoComplete(input string, options CallOptions) ([]Result, error)`**:
  receives an input string and call options and returns all possible results according to input string.
+ **`AutoCompleteWithContext(ctx context.Context, input string, options CallOptions) ([]Result, error)`**:
  same as `AutoComplete` but you can pass your context for more control.
+ **`Details(placeId string, options CallOptions) (Detail, error)`**:
  receives a `placeId` string and call options and returns details on that place id.
+ **`DetailsWithContext(ctx context.Context, placeId string, options CallOptions) (Detail, error)`**:
  same as `Details` but you can pass your context for more control.

### CallOptions

`CallOptions` is a struct that defines the behaviour of the operation. you can create a new `CallOptions`
with `search.NewDefaultCallOptions()`
function. you can customize the behaviour with passing multiple call options to the constructor.

list of call options for search service are:

+ [WithLocation(lat, lon float)](#): The geographic point that search would be around it. May be user current location
  or center of the city that you want search in it. If location is empty, search boundary is whole of the country.
+ [WithUserLocation(lat, lon float)](#): The current location of user. If you set location as center of a city, then you
  must fill this parameter with current location of user.
+ [WithEnglishLanguage()](#): sets the language for response to English. default is `fa` (Farsi)
+ [WithFarsiLanguage()](#): sets the language for response to Farsi. default is `fa` (Farsi)
+ [WithOriginRequestContext()](#): sets the context of current request to `origin`.
+ [WithFavouriteRequestContext()](#): sets the context of current request to `favouritre`.
+ [WithFirstDestinationRequestContext()](#): sets the context of current request to `destination1`.
+ [WithSecondDestinationRequestContext()](#): sets the context of current request to `destination2`.
+ [WithCityId(cityId int)](#): sets the city id for better results in search.
+ [WithHeaders(headers map[string]string)](#): sets custom headers for request.

This example will search for places like `Azadi` with city id of 1000 (Tehran) and around the given location:

```go
results, err := searchClient.AutoComplete("Azadi", search.NewDefaultCallOptions(
  search.WithCityId(1000),
  search.WithLocation(35.012, 53.1253),
))
if err != nil {
  panic(err)
}
```

## Area Gateways

After creating a config object, you can construct an area-gateways client for your services.

The constructor of area-gateways receives a config, version, timeout and multiple optional options.

some constructor options are:
+ [WithURL(url string)](#): could be used to override url of the service.
+ [WithTransport(transport *http.Transport)](#): could be used to set a new `http.Transport` for http client.
+ [WithRequestOpenTelemetryTracing(tracerName string)](#): could be used for activating opentelemetry tracing on client. for more info see [Here](#opentelemetry-tracing)

Example:

```go
package main

import (
	"fmt"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/config"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/services/area-gateways"
	"time"
)

func main() {
	cfg, err := config.NewDefaultConfig("api-key")
	if err != nil {
		panic(err)
	}

	areaGatewaysClient, err := area_gateways.NewAreaGatewaysClient(cfg, area_gateways.V1, time.Second)
	if err != nil {
		panic(err)
	}

	area, err := areaGatewaysClient.GetGateways(35.709374285391284, 51.40994310379028, area_gateways.NewDefaultCallOptions())
	if err != nil {
		panic(err)
	}

	fmt.Println(area)
}

```

### Operations

List of operations on a area-gateways client are:

+ **`GetGateways(lat, lon float64, options CallOptions) (Area, error)`**:
  it receives `lat`,`lon` as a location and CallOptions and returns a polygon and its Gateways. It will return an Empty
  area if no area is found with given lat and lon.
+ **`GetGatewaysWithContext(ctx context.Context, lat, lon float64, options CallOptions) (Area, error)`**:
  same as `GetGateways` but you can pass your context for more control.

### CallOptions

`CallOptions` is a struct that defines the behaviour of the operation. you can create a new `CallOptions`
with `area_gateways.NewDefaultCallOptions()`
function. you can customize the behaviour with passing multiple call options to the constructor.

list of call options for area-gateways service are:

+ [WithEnglishLanguage()](#): sets the language for response to English. default is `fa` (Farsi)
+ [WithFarsiLanguage()](#): sets the language for response to Farsi. default is `fa` (Farsi)
+ [WithHeaders(headers map[string]string)](#): sets custom headers for request.

This example will get gateways with farsi language with the given location:

```go
area, err := areaGatewaysClient.GetGateways(35.709374285391284, 51.40994310379028, area_gateways.NewDefaultCallOptions(
  area_gateways.WithFarsiLanguage()
))
if err != nil {
  panic(err)
}
```

## ETA

After creating a config object, you can construct an ETA client for your services.

The constructor of ETA client receives a config, version, timeout and multiple optional options.

some constructor options are:
+ [WithURL(url string)](#): could be used to override url of the service.
+ [WithTransport(transport *http.Transport)](#): could be used to set a new `http.Transport` for http client.
+ [WithRequestOpenTelemetryTracing(tracerName string)](#): could be used for activating opentelemetry tracing on client. for more info see [Here](#opentelemetry-tracing)

Example:

```go
package main

import (
	"fmt"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/config"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/services/eta"
	"time"
)

func main() {
	cfg, err := config.NewDefaultConfig("api-key")
	if err != nil {
		panic(err)
	}
	client, err := eta.NewETAClient(cfg, eta.V1, time.Second*10)
	if err != nil {
		panic(err)
	}

	results, err := client.GetETA([]eta.Point{
		{
			Lat: 35.77330981921435,
			Lon: 51.41834378242493,
		},
		{
			Lat: 35.739136559226864,
			Lon: 51.510804891586304,
		},
	}, eta.NewDefaultCallOptions(
		eta.WithHeaders(map[string]string{
			"foo": "bar",
		}),
		eta.WithNoTraffic(),
		WithEngine(EtaEngineV2),
	),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(results)
}


```

### Operations

List of operations on a eta client are:

+ **`GetETA(points []Point, options CallOptions) (ETA, error)`**:
  It will receive a list of point with minimum length of 2 and returns ETA. Will return error if less than 2 points are
  passed.
+ **`GetETAWithContext(ctx context.Context, points []Point, options CallOptions) (ETA, error)`**:
  same as `GetETA` but you can pass your context for more control.

### CallOptions

`CallOptions` is a struct that defines the behaviour of the operation. you can create a new `CallOptions`
with `eta.NewDefaultCallOptions()`
function. you can customize the behaviour with passing multiple call options to the constructor.

list of call options for eta service are:

+ [WithNoTraffic()](#): sets `no_traffic` query param ro true. with this option eta requests does not involve traffic
  data in response
+ [WithDepartureDateTime(dateTime string)](#): sets the departure date time of the eta request. (format of date time
  should be like `2006-01-02T15:04:05Z07:00` according to RFC3309)
+ [WithHeaders(headers map[string]string)](#): sets custom headers for request.

This example will get eta between the given locations:

```go
results, err := client.GetETA([]eta.Point{
    {
      Lat: 35.77330981921435,
      Lon: 51.41834378242493,
    },
    {
      Lat: 35.739136559226864,
      Lon: 51.510804891586304,
    },
  }, eta.NewDefaultCallOptions(
    eta.WithNoTraffic(),
  ),
)
if err != nil {
    panic(err)
}
```

## Matrix

After creating a config object, you can construct an Matrix client for your services.

The constructor of Matrix client receives a config, version, timeout and multiple optional options.

some constructor options are:
+ [WithURL(url string)](#): could be used to override url of the service.
+ [WithTransport(transport *http.Transport)](#): could be used to set a new `http.Transport` for http client.
+ [WithRequestOpenTelemetryTracing(tracerName string)](#): could be used for activating opentelemetry tracing on client. for more info see [Here](#opentelemetry-tracing)

Example:

```go
package main

import (
  "fmt"
  "gitlab.snapp.ir/Map/sdk/smapp-sdk-go/config"
  "gitlab.snapp.ir/Map/sdk/smapp-sdk-go/services/matrix"
  "time"
)

func main() {
  cfg, err := config.NewDefaultConfig("api-key")
  if err != nil {
    panic(err)
  }
  
  client, err := matrix.NewMatrixClient(cfg, matrix.V1, time.Second*10)
  if err != nil {
    panic(err)
  }

  result, err := client.GetMatrix([]matrix.Point{
    {
      Lat: 35.7733304928583,
      Lon: 51.418322660028934,
    },
    {
      Lat: 35.72895575080859,
      Lon: 51.37228488922119,
    },
  }, []matrix.Point{
    {
      Lat: 35.70033104179786,
      Lon: 51.351492404937744,
    },
    {
      Lat: 35.73933685292328,
      Lon: 51.50890588760376,
    },
  }, matrix.NewDefaultCallOptions(
	
	// if you want response with no traffic involved
	matrix.WithNoTraffic(),

	// if you want response with traffic involved
	// matrix.WithTraffic(),
	matrix.WithEngine(MatrixEngineV1),
    matrix.WithHeaders(map[string]string{
      "foo": "bar",
    }),
  ))

  if err != nil {
    panic(err)
  }

  fmt.Println(result)
}
```

### Operations

List of operations on a matrix client are:

+ **`GetMatrix(sources []Point, targets []Point, options CallOptions) (Output, error)`**:
  It will receive a list of points as sources and a list of points as targets and returns a matrix of eta predictions from all sources to all targets.
  Will return error if sources or targets are empty.
+ **`GetMatrixWithContext(ctx context.Context, sources []Point, targets []Point, options CallOptions) (Output, error)`**:
  same as `GetMatrix` but you can pass your context for more control.

### CallOptions

`CallOptions` is a struct that defines the behaviour of the operation. you can create a new `CallOptions`
with `matrix.NewDefaultCallOptions()`
function. you can customize the behaviour with passing multiple call options to the constructor.

list of call options for matrix service are:

+ [WithNoTraffic()](#): sets `no_traffic` query param ro true. with this option eta requests does not involve traffic data in response

+ [WithEngine(MatrixEngine)](#): sets the engine to be used in calculating the eta matrix.

+ [WithHeaders(headers map[string]string)](#): sets custom headers for request.

# Testing (Mocking)

There is a mock client for each service for usage of mocking. [mock](https://github.com/golang/mock) package is used for
this purpose.

an Example for mocking search client is:

```go
func TestFoo(t *testing.T) {
ctrl := gomock.NewController(t)

// Assert that Bar() is invoked.
defer ctrl.Finish()

mockSearch := search.NewMockSearchClient(ctrl)

}
```

you can use `mockSearch` to mock results of function calls of the client. for more information,
see [here](https://github.com/golang/mock)

# OpenTelemetry Tracing
`WithRequestOpenTelemetryTracing` could be used to enable [opentelemetry](https://opentelemetry.io/) on the sdk.

you should use functions with `WithContext` postfix of clients and pass a valid context of parent `Span`s for proper tracing. it will act unexpected if configurations of opentelemetry are not valid. 

> **Note** make sure you have declared a global tracer (with `SetTracerProvider` function existing in opentelemtry library) and proper propagator (`SetTextMapPropagator` function rxisting in opentelemetry library)

## Example
in this example we will configure a jaeger opentelemetry client with `area-gateways` client.

```go
package main

import (
	"context"
	"fmt"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/config"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/services/area-gateways"
	jaeger_propagator "go.opentelemetry.io/contrib/propagators/jaeger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	traceSpan "go.opentelemetry.io/otel/trace"
	"time"
)

func main() {
	exp, err := jaeger.New(jaeger.WithAgentEndpoint(
			jaeger.WithAgentHost("localhost"),
			jaeger.WithAgentPort("6831"),
		),
	)
	if err != nil {
		panic(err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithSampler(trace.ParentBased(trace.TraceIDRatioBased(0.5))),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("sample-tracer"),
		),
		),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(jaeger_propagator.Jaeger{})

	cfg, err := config.NewDefaultConfig("api-key")
	if err != nil {
		panic(err)
	}

	areaGatewaysClient, err := area_gateways.NewAreaGatewaysClient(cfg, area_gateways.V1, time.Second,
		area_gateways.WithRequestOpenTelemetryTracing("sample-tracer"),
	)
	if err != nil {
		panic(err)
	}

	var sp traceSpan.Span
	ctx, sp := otel.Tracer("test-test").Start(context.Background(), "start")
	area, err := areaGatewaysClient.GetGatewaysWithContext(ctx, 35.709374285391284, 51.40994310379028, area_gateways.NewDefaultCallOptions())
	if err != nil {
		panic(err)
	}

	fmt.Println(area)
	sp.End()
	time.Sleep(10 * time.Second)
}

```
