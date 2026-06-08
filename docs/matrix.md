# Matrix

After creating a config object, construct a Matrix client.

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
	"github.com/snapp-incubator/smapp-sdk-go/services/matrix"
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

	result, err := client.GetMatrix(
		[]matrix.Point{
			{Lat: 35.7733304928583, Lon: 51.418322660028934},
			{Lat: 35.72895575080859, Lon: 51.37228488922119},
		},
		[]matrix.Point{
			{Lat: 35.70033104179786, Lon: 51.351492404937744},
			{Lat: 35.73933685292328, Lon: 51.50890588760376},
		},
		matrix.NewDefaultCallOptions(
			matrix.WithNoTraffic(),
		),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
```

## Operations

- `GetMatrix(sources []Point, targets []Point, options CallOptions) (Output, error)` — returns ETA matrix from all sources to all targets; error if sources or targets are empty
- `GetMatrixWithContext(ctx context.Context, sources []Point, targets []Point, options CallOptions) (Output, error)`

## CallOptions

Create with `matrix.NewDefaultCallOptions()`.

| Option | Description |
|---|---|
| `WithNoTraffic()` | Exclude traffic data |
| `WithTraffic()` | Include traffic data |
| `WithEngine(MatrixEngine)` | Set calculation engine |
| `WithHeaders(map[string]string)` | Custom request headers |
