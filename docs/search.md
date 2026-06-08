# Search

After creating a config object, construct a search client.

Constructor options:
- `WithURL(url string)` — override service URL
- `WithTransport(transport http.RoundTripper)` — set custom HTTP transport
- `WithRequestOpenTelemetryTracing(tracerName string)` — enable OpenTelemetry tracing ([details](opentelemetry.md))

## Example

```go
package main

import (
	"fmt"
	"github.com/snapp-incubator/smapp-sdk-go/config"
	"github.com/snapp-incubator/smapp-sdk-go/services/search"
	"time"
)

func main() {
	cfg, err := config.NewDefaultConfig("api-key")
	if err != nil {
		panic(err)
	}

	searchClient, err := search.NewSearchClient(cfg, search.V1, time.Second,
		search.WithURL("https://new-url.com"), // optional
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

## Operations

- `GetCities(options CallOptions) ([]City, error)`
- `GetCitiesWithContext(ctx context.Context, options CallOptions) ([]City, error)`
- `SearchCity(input string, options CallOptions) ([]City, error)`
- `SearchCityWithContext(ctx context.Context, input string, options CallOptions) ([]City, error)`
- `AutoComplete(input string, options CallOptions) ([]Result, error)`
- `AutoCompleteWithContext(ctx context.Context, input string, options CallOptions) ([]Result, error)`
- `Details(placeId string, options CallOptions) (Detail, error)`
- `DetailsWithContext(ctx context.Context, placeId string, options CallOptions) (Detail, error)`

## CallOptions

Create with `search.NewDefaultCallOptions()`.

| Option | Description |
|---|---|
| `WithLocation(lat, lon float)` | Search center point |
| `WithUserLocation(lat, lon float)` | Current user location |
| `WithEnglishLanguage()` | Response in English |
| `WithFarsiLanguage()` | Response in Farsi (default) |
| `WithOriginRequestContext()` | Request context: `origin` |
| `WithFavouriteRequestContext()` | Request context: `favourite` |
| `WithFirstDestinationRequestContext()` | Request context: `destination1` |
| `WithSecondDestinationRequestContext()` | Request context: `destination2` |
| `WithCityId(int)` | City ID for better results |
| `WithHeaders(map[string]string)` | Custom request headers |

```go
results, err := searchClient.AutoComplete("Azadi", search.NewDefaultCallOptions(
	search.WithCityId(1000),
	search.WithLocation(35.012, 53.1253),
))
```
