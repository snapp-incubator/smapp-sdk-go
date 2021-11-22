package area_gateways

import (
	"context"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/config"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewAreaGatewaysClient(t *testing.T) {
	t.Run("without_options", func(t *testing.T) {
		cfg, err := config.NewDefaultConfig("key")
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewAreaGatewaysClient(cfg, V1, time.Second)
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
		client, err := NewAreaGatewaysClient(cfg, V1, time.Second, WithURL("https://google.com"))
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

func TestClient_GetGateways(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"id":"460353260","name":"بیمارستان فیروزگر","type":"Polygon","coordinates":[[[51.41066998243332,35.71003854872514],[51.40912737697363,35.709946260011115],[51.40923164784908,35.70876228705955],[51.41077123582363,35.70885212696658],[51.41066998243332,35.71003854872514]]],"gates":[{"name":"درِ درمانگاه","type":"Point","coordinates":[51.41055263578892,35.71002711473095]},{"name":"درِ اورژانس","type":"Point","coordinates":[51.41074977815151,35.70907046474276]},{"name":"درِ اصلی","type":"Point","coordinates":[51.4106926,35.7097484]}]}`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.HeaderSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewAreaGatewaysClient(cfg, V1, time.Millisecond*100, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		_, err = client.GetGateways(35.709374285391284, 51.40994310379028, NewDefaultCallOptions(
			WithFarsiLanguage(),
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
		client, err := NewAreaGatewaysClient(cfg, V1, time.Millisecond*100, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		_, err = client.GetGateways(35.709374285391284, 51.40994310379028, NewDefaultCallOptions(
			WithFarsiLanguage(),
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
		client, err := NewAreaGatewaysClient(cfg, V1, time.Millisecond*100, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		_, err = client.GetGateways(35.709374285391284, 51.40994310379028, NewDefaultCallOptions(
			WithFarsiLanguage(),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("should not be nil because response is invalid")
		}
	})
	t.Run("invalid_input", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{}`))
		}))

		cfg, err := config.NewDefaultConfig("key")
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewAreaGatewaysClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		_, err = client.GetGateways(5000, 5000, NewDefaultCallOptions())
		if err == nil {
			t.Fatalf("err should not be nil due to invalid input lat and lon")
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

		client, err := NewAreaGatewaysClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		_, err = client.GetGateways(35.77331417156089, 51.41831696033478, NewDefaultCallOptions(
			WithFarsiLanguage(),
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
		client, err := NewAreaGatewaysClient(cfg, V1, time.Millisecond*100, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		_, err = client.GetGateways(35.77331417156089, 51.41831696033478, NewDefaultCallOptions(
			WithFarsiLanguage(),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("there should be an errordue to timeout")
		}
	})
}

func TestClient_GetGatewaysWithContext(t *testing.T) {
	t.Run("invalid_request", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{}`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.QueryParamSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewAreaGatewaysClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		var ctx context.Context = nil
		_, err = client.GetGatewaysWithContext(ctx, 35.77331417156089, 51.41831696033478, NewDefaultCallOptions(
			WithFarsiLanguage(),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("there should be an error when creating request")
		}
	})
}
