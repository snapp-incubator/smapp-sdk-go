package eta

import (
	"net/http"
	"testing"
	"time"

	"github.com/snapp-incubator/smapp-sdk-go/config"
)

func TestWithURL(t *testing.T) {
	cfg, err := config.NewDefaultConfig("key")
	if err != nil {
		t.Fatalf("could not create default config due to: %s", err.Error())
	}
	client, err := NewETAClient(cfg, V1, time.Second, WithURL("https://google.com"))
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
	client, err := NewETAClient(cfg, V1, time.Second, WithTransport(&http.Transport{
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
	client, err := NewETAClient(cfg, V1, time.Second,
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

func TestWithPathStyleLegacy(t *testing.T) {
	cfg, err := config.NewDefaultConfig("key")
	if err != nil {
		t.Fatalf("could not create default config due to: %s", err.Error())
	}
	cfg.APIBaseURL = "https://api.example.com"

	client, err := NewETAClient(cfg, V1, time.Second, WithPathStyle(LegacyPathStyle))
	if err != nil {
		t.Fatalf("could not create eta client due to: %s", err.Error())
	}

	expectedURL := "https://api.example.com/eta/v1"
	if client.url != expectedURL {
		t.Fatalf("client.url should be %s but it is %s", expectedURL, client.url)
	}

	if client.pathStyle != LegacyPathStyle {
		t.Fatalf("client.pathStyle should be %d but it is %d", LegacyPathStyle, client.pathStyle)
	}
}

func TestWithPathStyleV1(t *testing.T) {
	cfg, err := config.NewDefaultConfig("key")
	if err != nil {
		t.Fatalf("could not create default config due to: %s", err.Error())
	}
	cfg.APIBaseURL = "https://api.example.com"

	client, err := NewETAClient(cfg, V1, time.Second, WithPathStyle(V1PathStyle))
	if err != nil {
		t.Fatalf("could not create eta client due to: %s", err.Error())
	}

	expectedURL := "https://api.example.com/api/v1/eta"
	if client.url != expectedURL {
		t.Fatalf("client.url should be %s but it is %s", expectedURL, client.url)
	}

	if client.pathStyle != V1PathStyle {
		t.Fatalf("client.pathStyle should be %d but it is %d", V1PathStyle, client.pathStyle)
	}
}

func TestDefaultPathStyle(t *testing.T) {
	cfg, err := config.NewDefaultConfig("key")
	if err != nil {
		t.Fatalf("could not create default config due to: %s", err.Error())
	}
	cfg.APIBaseURL = "https://api.example.com"

	client, err := NewETAClient(cfg, V1, time.Second)
	if err != nil {
		t.Fatalf("could not create eta client due to: %s", err.Error())
	}

	// Default should be LegacyPathStyle
	expectedURL := "https://api.example.com/eta/v1"
	if client.url != expectedURL {
		t.Fatalf("client.url should be %s but it is %s", expectedURL, client.url)
	}

	if client.pathStyle != LegacyPathStyle {
		t.Fatalf("client.pathStyle should be %d (LegacyPathStyle) but it is %d", LegacyPathStyle, client.pathStyle)
	}
}
