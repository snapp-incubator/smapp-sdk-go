# OpenTelemetry Tracing

Pass `WithRequestOpenTelemetryTracing(tracerName string)` to any service constructor to enable [OpenTelemetry](https://opentelemetry.io/) tracing.

Use the `WithContext` variants of client methods and pass a context containing a valid parent span. Requires a global tracer (`otel.SetTracerProvider`) and propagator (`otel.SetTextMapPropagator`) to be configured before use.

## Example (Jaeger + area-gateways)

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/snapp-incubator/smapp-sdk-go/config"
	area_gateways "github.com/snapp-incubator/smapp-sdk-go/services/area-gateways"
	jaeger_propagator "go.opentelemetry.io/contrib/propagators/jaeger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	traceSpan "go.opentelemetry.io/otel/trace"
)

func main() {
	exp, err := jaeger.New(jaeger.WithAgentEndpoint(
		jaeger.WithAgentHost("localhost"),
		jaeger.WithAgentPort("6831"),
	))
	if err != nil {
		panic(err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithSampler(trace.ParentBased(trace.TraceIDRatioBased(0.5))),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("sample-tracer"),
		)),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(jaeger_propagator.Jaeger{})

	cfg, err := config.NewDefaultConfig("api-key")
	if err != nil {
		panic(err)
	}

	client, err := area_gateways.NewAreaGatewaysClient(cfg, area_gateways.V1, time.Second,
		area_gateways.WithRequestOpenTelemetryTracing("sample-tracer"),
	)
	if err != nil {
		panic(err)
	}

	var sp traceSpan.Span
	ctx, sp := otel.Tracer("test-test").Start(context.Background(), "start")
	defer sp.End()

	area, err := client.GetGatewaysWithContext(ctx, 35.709374285391284, 51.40994310379028, area_gateways.NewDefaultCallOptions())
	if err != nil {
		panic(err)
	}

	fmt.Println(area)
	time.Sleep(10 * time.Second) // wait for spans to flush
}
```
