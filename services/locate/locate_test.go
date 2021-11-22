package locate

import (
	"context"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/config"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClient_LocatePoints(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`[{"input":{"lat":35.70974,"lon":51.408699},"snapped_points":[{"point":{"lat":35.70998,"lon":51.40868}}]}]`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.HeaderSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewLocateClient(cfg, V1, time.Millisecond*100, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		_, err = client.LocatePoints([]Point{{
			Lat: 35.70973799747619,
			Lon: 51.40869855880737,
		}}, NewDefaultCallOptions(
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err != nil {
			t.Fatalf("should be nil.")
		}
	})
	t.Run("non-200-status", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.HeaderSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewLocateClient(cfg, V1, time.Millisecond*100, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		_, err = client.LocatePoints([]Point{{
			Lat: 35.70973799747619,
			Lon: 51.40869855880737,
		}}, NewDefaultCallOptions(
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("should not be nil.")
		}
	})
	t.Run("invalid_response", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.HeaderSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewLocateClient(cfg, V1, time.Millisecond*100, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		_, err = client.LocatePoints([]Point{{
			Lat: 35.70973799747619,
			Lon: 51.40869855880737,
		}}, NewDefaultCallOptions(
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("should not be nil because response is invalid")
		}
	})
	t.Run("invalid_apikey_source", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{}`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.QueryParamSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		cfg.APIKeySource = "foo"

		client, err := NewLocateClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		_, err = client.LocatePoints([]Point{{
			Lat: 35.70973799747619,
			Lon: 51.40869855880737,
		}}, NewDefaultCallOptions(
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("there should be an error with api key source")
		}
	})
	t.Run("timeout", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(200 * time.Millisecond)
			_, _ = w.Write([]byte(`{}`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.QueryParamSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewLocateClient(cfg, V1, time.Millisecond*100, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		_, err = client.LocatePoints([]Point{{
			Lat: 35.70973799747619,
			Lon: 51.40869855880737,
		}}, NewDefaultCallOptions(
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("there should be an errordue to timeout")
		}
	})
}

func TestClient_LocatePointsWithContext(t *testing.T) {
	t.Run("invalid_request", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{}`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.QueryParamSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewLocateClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		var ctx context.Context = nil
		_, err = client.LocatePointsWithContext(ctx, []Point{{
			Lat: 35.70973799747619,
			Lon: 51.40869855880737,
		}}, NewDefaultCallOptions(
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))

		if err == nil {
			t.Fatalf("there should be an error when creating request")
		}
	})
}

func TestNewLocateClient(t *testing.T) {
	t.Run("without_options", func(t *testing.T) {
		cfg, err := config.NewDefaultConfig("key")
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewLocateClient(cfg, V1, time.Second)
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		if client == nil {
			t.Fatalf("client should not be nil")
		}
	})
	t.Run("with_options", func(t *testing.T) {
		cfg, err := config.NewDefaultConfig("key")
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewLocateClient(cfg, V1, time.Second, WithURL("https://google.com"))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		if client == nil {
			t.Fatalf("client should not be nil")
		}

		if client.url != "https://google.com" {
			t.Fatalf("client.url should be %s but it is %s", "https://google.com", client.url)
		}
	})
}
