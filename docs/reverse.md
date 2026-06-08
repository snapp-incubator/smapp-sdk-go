# Reverse Geocode

After creating a config object, construct a reverse geocode client.

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
	"github.com/snapp-incubator/smapp-sdk-go/services/reverse"
	"time"
)

func main() {
	cfg, err := config.NewDefaultConfig("api-key")
	if err != nil {
		panic(err)
	}

	reverseClient, err := reverse.NewReverseClient(cfg, reverse.V1, time.Second,
		reverse.WithURL("https://new-url.com"), // optional
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

## Operations

- `GetComponents(lat, lon float64, options CallOptions) ([]Component, error)`
- `GetComponentsWithContext(ctx context.Context, lat, lon float64, options CallOptions) ([]Component, error)`
- `GetDisplayName(lat, lon float64, options CallOptions) (string, error)`
- `GetDisplayNameWithContext(ctx context.Context, lat, lon float64, options CallOptions) (string, error)`
- `GetFrequent(lat, lon float64, options CallOptions) (FrequentAddress, error)`
- `GetFrequentWithContext(ctx context.Context, lat, lon float64, options CallOptions) (FrequentAddress, error)`
- `GetBatch(request BatchReverseRequest) ([]Result, error)`
- `GetBatchWithContext(ctx context.Context, request BatchReverseRequest) ([]Result, error)`
- `GetBatchDisplayName(request BatchReverseRequest) ([]Result, error)`
- `GetBatchDisplayNameWithContext(ctx context.Context, request BatchReverseRequest) ([]Result, error)`
- `GetStructuralResult(lat, lon float64, options CallOptions) ([]StructuralComponent, error)`
- `GetStructuralResultWithContext(ctx context.Context, lat, lon float64, options CallOptions) (*StructuralComponent, error)`
- `GetBatchStructuralResults(request BatchReverseRequest) ([]StructuralResult, error)`
- `GetBatchStructuralResultsWithContext(request BatchReverseRequest) ([]StructuralResult, error)`

> You can iterate over `StructuralComponent` fields using `NewIterator()`.

## CallOptions

Create with `reverse.NewDefaultCallOptions()`.

| Option | Description |
|---|---|
| `WithZoomLevel(int)` | Zoom level, default `16` |
| `WithEnglishLanguage()` | Response in English |
| `WithFarsiLanguage()` | Response in Farsi (default) |
| `WithArabicLanguage()` | Response in Arabic |
| `WithPassengerResponseType()` | Response type for passengers |
| `WithDriverResponseType()` | Response type for drivers |
| `WithVerboseResponseType()` | Verbose response |
| `WithBikerResponseType()` | Response type for bikers |
| `WithOriginResponseType()` | Response type for origin points |
| `WithDestinationResponseType()` | Response type for destination points |
| `WithIraqResponseType()` | Response type for Iraq (Baly only) |
| `WithHeaders(map[string]string)` | Custom request headers |

```go
displayName, err := reverseClient.GetDisplayName(35.0123, 53.12312, reverse.NewDefaultCallOptions(
	reverse.WithZoomLevel(17),
	reverse.WithEnglishLanguage(),
))
```
