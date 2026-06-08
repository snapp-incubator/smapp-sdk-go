# Area Gateways

After creating a config object, construct an area-gateways client.

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
	area_gateways "github.com/snapp-incubator/smapp-sdk-go/services/area-gateways"
	"time"
)

func main() {
	cfg, err := config.NewDefaultConfig("api-key")
	if err != nil {
		panic(err)
	}

	client, err := area_gateways.NewAreaGatewaysClient(cfg, area_gateways.V1, time.Second)
	if err != nil {
		panic(err)
	}

	area, err := client.GetGateways(35.709374285391284, 51.40994310379028, area_gateways.NewDefaultCallOptions())
	if err != nil {
		panic(err)
	}

	fmt.Println(area)
}
```

## Operations

- `GetGateways(lat, lon float64, options CallOptions) (Area, error)` — returns polygon and its gateways; empty area if none found
- `GetGatewaysWithContext(ctx context.Context, lat, lon float64, options CallOptions) (Area, error)`

## CallOptions

Create with `area_gateways.NewDefaultCallOptions()`.

| Option | Description |
|---|---|
| `WithEnglishLanguage()` | Response in English |
| `WithFarsiLanguage()` | Response in Farsi (default) |
| `WithHeaders(map[string]string)` | Custom request headers |

```go
area, err := client.GetGateways(35.709374285391284, 51.40994310379028, area_gateways.NewDefaultCallOptions(
	area_gateways.WithFarsiLanguage(),
))
```
