# ETA

After creating a config object, construct an ETA client.

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
	"github.com/snapp-incubator/smapp-sdk-go/services/eta"
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
		{Lat: 35.77330981921435, Lon: 51.41834378242493},
		{Lat: 35.739136559226864, Lon: 51.510804891586304},
	}, eta.NewDefaultCallOptions(
		eta.WithNoTraffic(),
	))
	if err != nil {
		panic(err)
	}

	fmt.Println(results)
}
```

## Operations

- `GetETA(points []Point, options CallOptions) (ETA, error)` — minimum 2 points required
- `GetETAWithContext(ctx context.Context, points []Point, options CallOptions) (ETA, error)`

## CallOptions

Create with `eta.NewDefaultCallOptions()`.

| Option | Description |
|---|---|
| `WithNoTraffic()` | Exclude traffic data |
| `WithDepartureDateTime(string)` | Departure time (RFC3339: `2006-01-02T15:04:05Z07:00`) |
| `WithHeaders(map[string]string)` | Custom request headers |
