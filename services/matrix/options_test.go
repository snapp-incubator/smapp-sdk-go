package matrix

import (
	"github.com/snapp-incubator/smapp-sdk-go/config"
	"net/http"
	"testing"
	"time"
)

func TestWithURL(t *testing.T) {
	cfg, err := config.NewDefaultConfig("key")
	if err != nil {
		t.Fatalf("could not create default config due to: %s", err.Error())
	}
	client, err := NewMatrixClient(cfg, V1, time.Second, WithURL("https://google.com"))
	if err != nil {
		t.Fatalf("could not create reverse client due to: %s", err.Error())
	}

	if client.url != "https://google.com" {
		t.Fatalf("client.URL should be %s but it is %s", "https://google.com", client.url)
	}
}

func TestWithTransport(t *testing.T) {
	cfg, err := config.NewDefaultConfig("key")
	if err != nil {
		t.Fatalf("could not create default config due to: %s", err.Error())
	}
	client, err := NewMatrixClient(cfg, V1, time.Second, WithTransport(&http.Transport{
		MaxIdleConns: 2,
	}))
	if err != nil {
		t.Fatalf("could not create search client due to: %s", err.Error())
	}

	if client.httpClient.Transport.(*http.Transport).MaxIdleConns != 2 {
		t.Fatalf("client.httpClient.Transport.MaxIdleConns should be %d but it is %d", 2, client.httpClient.Transport.(*http.Transport).MaxIdleConns)
	}
}

func TestWithRequestOpenTelemetryTracing(t *testing.T) {
	cfg, err := config.NewDefaultConfig("key")
	if err != nil {
		t.Fatalf("could not create default config due to: %s", err.Error())
	}
	client, err := NewMatrixClient(cfg, V1, time.Second,
		WithTransport(&http.Transport{
			MaxIdleConns: 2,
		}),
		WithRequestOpenTelemetryTracing("test"),
	)
	if err != nil {
		t.Fatalf("could not create search client due to: %s", err.Error())
	}

	if client.tracerName != "test" {
		t.Fatalf("client.tracerName should be %s but it is %s", "test", client.tracerName)
	}
}

func TestWithClient(t *testing.T) {
	cfg, err := config.NewDefaultConfig("key")
	if err != nil {
		t.Fatalf("could not create default config due to: %s", err.Error())
	}
	client, err := NewMatrixClient(cfg, V1, time.Second, WithClient(&http.Client{
		Timeout: time.Second * 10,
	}))
	if err != nil {
		t.Fatalf("could not create search client due to: %s", err.Error())
	}

	if client.httpClient.Timeout.Seconds() != 10 {
		t.Fatalf("client.httpClient.Timeout should be %d but it is %d", 10, client.httpClient.Timeout)
	}
}
