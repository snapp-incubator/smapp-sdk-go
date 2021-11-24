package eta

import (
	"context"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/config"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewETAClient(t *testing.T) {
	t.Run("without_options", func(t *testing.T) {
		cfg, err := config.NewDefaultConfig("key")
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewETAClient(cfg, V1, time.Second)
		if err != nil {
			t.Fatalf("could not create eta client due to: %s", err.Error())
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
		client, err := NewETAClient(cfg, V1, time.Second, WithURL("https://google.com"))
		if err != nil {
			t.Fatalf("could not create eta client due to: %s", err.Error())
		}
		if client == nil {
			t.Fatalf("client should not be nil")
		}

		if client.url != "https://google.com" {
			t.Fatalf("client.url should be %s but it is %s", "https://google.com", client.url)
		}
	})
}

func TestClient_GetETA(t *testing.T) {
	t.Run("non-200-status", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.HeaderSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewETAClient(cfg, V1, time.Millisecond*100, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create locate client due to: %s", err.Error())
		}
		_, err = client.GetETA([]Point{
			{
				Lat: 35.70973799747619,
				Lon: 51.40869855880737,
			},
			{
				Lat: 35.70973799747619,
				Lon: 51.40969855880737,
			},
		}, NewDefaultCallOptions(
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
		client, err := NewETAClient(cfg, V1, time.Millisecond*100, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create locate client due to: %s", err.Error())
		}
		_, err = client.GetETA([]Point{
			{
				Lat: 35.70973799747619,
				Lon: 51.40869855880737,
			},
			{
				Lat: 35.70973799747619,
				Lon: 51.40969855880737,
			},
		}, NewDefaultCallOptions(
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("should not be nil because response is invalid")
		}
	})

	t.Run("invalid_input_points", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.HeaderSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewETAClient(cfg, V1, time.Millisecond*100, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create locate client due to: %s", err.Error())
		}
		_, err = client.GetETA([]Point{
			{
				Lat: 35.70973799747619,
				Lon: 51.40869855880737,
			},
		}, NewDefaultCallOptions(
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("should not be nil because input points are invalid")
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

		client, err := NewETAClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create locate client due to: %s", err.Error())
		}
		_, err = client.GetETA([]Point{
			{
				Lat: 35.70973799747619,
				Lon: 51.40869855880737,
			},
			{
				Lat: 35.70973799747619,
				Lon: 51.40969855880737,
			},
		}, NewDefaultCallOptions(
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
		client, err := NewETAClient(cfg, V1, time.Millisecond*100, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create locate client due to: %s", err.Error())
		}
		_, err = client.GetETA([]Point{
			{
				Lat: 35.70973799747619,
				Lon: 51.40869855880737,
			},
			{
				Lat: 35.70973799747619,
				Lon: 51.40969855880737,
			},
		}, NewDefaultCallOptions(
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("there should be an errordue to timeout")
		}
	})

	t.Run("valid", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"trip":{"legs":[{"time": 180,"length":3900}]}}`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.QueryParamSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewETAClient(cfg, V1, time.Millisecond*100, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create locate client due to: %s", err.Error())
		}
		_, err = client.GetETA([]Point{
			{
				Lat: 35.70973799747619,
				Lon: 51.40869855880737,
			},
			{
				Lat: 35.70973799747619,
				Lon: 51.40969855880737,
			},
		}, NewDefaultCallOptions(
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
			WithNoTraffic(),
			WithDepartureDateTime("2020-01-1"),
		))
		if err != nil {
			t.Fatalf("there should not be an error becuase request is valid")
		}
	})
}

func TestClient_GetETAWithContext(t *testing.T) {
	t.Run("invalid_request", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{}`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.QueryParamSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewETAClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create eta client due to: %s", err.Error())
		}
		var ctx context.Context = nil
		_, err = client.GetETAWithContext(ctx, []Point{
			{
				Lat: 35.70973799747619,
				Lon: 51.40869855880737,
			},
			{
				Lat: 35.70973799747619,
				Lon: 51.41069855880737,
			},
		}, NewDefaultCallOptions(
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))

		if err == nil {
			t.Fatalf("there should be an error when creating request")
		}
	})
}
