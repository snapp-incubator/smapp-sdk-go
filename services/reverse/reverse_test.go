package reverse

import (
	"context"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/config"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewReverseClient(t *testing.T) {
	t.Run("without_options", func(t *testing.T) {
		cfg, err := config.NewDefaultConfig("key")
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewReverseClient(cfg, V1, time.Second)
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
		client, err := NewReverseClient(cfg, V1, time.Second, WithURL("https://google.com"))
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

func TestClient_GetComponents(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"status":"OK","result":{"components":[{"name":"اسنپ","type":"company","distance":0},{"name":"تقاطع مهرداد","type":"relation","distance":3},{"name":"سید رضا سعیدی","type":"residential","distance":3},{"name":"جردن - پارک ملت","type":"meta_neighbourhood"},{"name":"تهران","type":"meta_city"}]},"traffic_zone":{"in_central":false,"in_evenodd":false}}`))
		}))

		cfg, err := config.NewDefaultConfig("key")
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewReverseClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		components, err := client.GetComponents(35.77331417156089, 51.41831696033478, NewDefaultCallOptions(
			WithZoomLevel(17),
			WithFarsiLanguage(),
			WithPassengerResponseType(),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err != nil {
			t.Fatalf("could not get components: %s", err.Error())
		}
		if len(components) != 5 {
			t.Fatalf("there should be 5 components")
		}
	})
	t.Run("invalid_response", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.QueryParamSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewReverseClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		_, err = client.GetComponents(35.77331417156089, 51.41831696033478, NewDefaultCallOptions(
			WithZoomLevel(17),
			WithFarsiLanguage(),
			WithPassengerResponseType(),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("there should be an error when parsing request")
		}
	})
	t.Run("invalid_apikey_source", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"status":"OK","result":{"components":[{"name":"اسنپ","type":"company","distance":0},{"name":"تقاطع مهرداد","type":"relation","distance":3},{"name":"سید رضا سعیدی","type":"residential","distance":3},{"name":"جردن - پارک ملت","type":"meta_neighbourhood"},{"name":"تهران","type":"meta_city"}]},"traffic_zone":{"in_central":false,"in_evenodd":false}}`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.QueryParamSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		cfg.APIKeySource = "foo"

		client, err := NewReverseClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		_, err = client.GetComponents(35.77331417156089, 51.41831696033478, NewDefaultCallOptions(
			WithZoomLevel(17),
			WithFarsiLanguage(),
			WithPassengerResponseType(),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("there should be an error with api key source")
		}
	})
	t.Run("error_status", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"status": "ERROR"}`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.QueryParamSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewReverseClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		_, err = client.GetComponents(35.77331417156089, 51.41831696033478, NewDefaultCallOptions(
			WithZoomLevel(17),
			WithFarsiLanguage(),
			WithPassengerResponseType(),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("there should be an error. status is ERROR")
		}
	})
	t.Run("non_200_status", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
			_, _ = w.Write([]byte(`{"status":"OK","result":{"components":[{"name":"اسنپ","type":"company","distance":0},{"name":"تقاطع مهرداد","type":"relation","distance":3},{"name":"سید رضا سعیدی","type":"residential","distance":3},{"name":"جردن - پارک ملت","type":"meta_neighbourhood"},{"name":"تهران","type":"meta_city"}]},"traffic_zone":{"in_central":false,"in_evenodd":false}}`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.QueryParamSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewReverseClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		_, err = client.GetComponents(35.77331417156089, 51.41831696033478, NewDefaultCallOptions(
			WithZoomLevel(17),
			WithFarsiLanguage(),
			WithPassengerResponseType(),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("there should be an error. status is 500")
		}
	})
	t.Run("timeout", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(200 * time.Millisecond)
			_, _ = w.Write([]byte(`{"status":"OK","result":{"components":[{"name":"اسنپ","type":"company","distance":0},{"name":"تقاطع مهرداد","type":"relation","distance":3},{"name":"سید رضا سعیدی","type":"residential","distance":3},{"name":"جردن - پارک ملت","type":"meta_neighbourhood"},{"name":"تهران","type":"meta_city"}]},"traffic_zone":{"in_central":false,"in_evenodd":false}}`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.QueryParamSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewReverseClient(cfg, V1, time.Millisecond*100, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		_, err = client.GetComponents(35.77331417156089, 51.41831696033478, NewDefaultCallOptions(
			WithZoomLevel(17),
			WithFarsiLanguage(),
			WithPassengerResponseType(),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("there should be an errordue to timeout")
		}
	})
}

func TestClient_GetComponentsWithContext(t *testing.T) {
	t.Run("invalid_request", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{}`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.QueryParamSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewReverseClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		var ctx context.Context = nil
		_, err = client.GetComponentsWithContext(ctx, 35.77331417156089, 51.41831696033478, NewDefaultCallOptions(
			WithZoomLevel(17),
			WithFarsiLanguage(),
			WithPassengerResponseType(),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("there should be an error when creating request")
		}
	})
}

func TestClient_GetDisplayName(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"status":"OK","result":{"displayName":"اسنپ"},"traffic_zone":{"in_central":false,"in_evenodd":false}}`))
		}))

		cfg, err := config.NewDefaultConfig("key")
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewReverseClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		displayName, err := client.GetDisplayName(35.77331417156089, 51.41831696033478, NewDefaultCallOptions(
			WithZoomLevel(17),
			WithFarsiLanguage(),
			WithPassengerResponseType(),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err != nil {
			t.Fatalf("could not get components: %s", err.Error())
		}
		if displayName != "اسنپ" {
			t.Fatalf("invalid_address")
		}
	})
	t.Run("invalid_response", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.QueryParamSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewReverseClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		_, err = client.GetDisplayName(35.77331417156089, 51.41831696033478, NewDefaultCallOptions(
			WithZoomLevel(17),
			WithFarsiLanguage(),
			WithPassengerResponseType(),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("there should be an error when parsing request")
		}
	})
	t.Run("invalid_apikey_source", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"status":"OK","result":{"components":[{"name":"اسنپ","type":"company","distance":0},{"name":"تقاطع مهرداد","type":"relation","distance":3},{"name":"سید رضا سعیدی","type":"residential","distance":3},{"name":"جردن - پارک ملت","type":"meta_neighbourhood"},{"name":"تهران","type":"meta_city"}]},"traffic_zone":{"in_central":false,"in_evenodd":false}}`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.QueryParamSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		cfg.APIKeySource = "foo"

		client, err := NewReverseClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		_, err = client.GetDisplayName(35.77331417156089, 51.41831696033478, NewDefaultCallOptions(
			WithZoomLevel(17),
			WithFarsiLanguage(),
			WithPassengerResponseType(),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("there should be an error with apikey source")
		}
	})
	t.Run("error_status", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"status": "ERROR"}`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.QueryParamSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewReverseClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		_, err = client.GetDisplayName(35.77331417156089, 51.41831696033478, NewDefaultCallOptions(
			WithZoomLevel(17),
			WithFarsiLanguage(),
			WithPassengerResponseType(),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("there should be an error. status is ERROR")
		}
	})
	t.Run("non_200_status", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
			_, _ = w.Write([]byte(`{"status":"OK","result":{"components":[{"name":"اسنپ","type":"company","distance":0},{"name":"تقاطع مهرداد","type":"relation","distance":3},{"name":"سید رضا سعیدی","type":"residential","distance":3},{"name":"جردن - پارک ملت","type":"meta_neighbourhood"},{"name":"تهران","type":"meta_city"}]},"traffic_zone":{"in_central":false,"in_evenodd":false}}`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.QueryParamSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewReverseClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		_, err = client.GetDisplayName(35.77331417156089, 51.41831696033478, NewDefaultCallOptions(
			WithZoomLevel(17),
			WithFarsiLanguage(),
			WithPassengerResponseType(),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("there should be an error. status is 500")
		}
	})
	t.Run("timeout", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(200 * time.Millisecond)
			_, _ = w.Write([]byte(`{"status":"OK","result":{"components":[{"name":"اسنپ","type":"company","distance":0},{"name":"تقاطع مهرداد","type":"relation","distance":3},{"name":"سید رضا سعیدی","type":"residential","distance":3},{"name":"جردن - پارک ملت","type":"meta_neighbourhood"},{"name":"تهران","type":"meta_city"}]},"traffic_zone":{"in_central":false,"in_evenodd":false}}`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.QueryParamSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewReverseClient(cfg, V1, time.Millisecond*100, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		_, err = client.GetDisplayName(35.77331417156089, 51.41831696033478, NewDefaultCallOptions(
			WithZoomLevel(17),
			WithFarsiLanguage(),
			WithPassengerResponseType(),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("there should be an errordue to timeout")
		}
	})
}

func TestClient_GetDisplayNameWithContext(t *testing.T) {
	t.Run("invalid_request", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{}`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.QueryParamSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewReverseClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create reverse client due to: %s", err.Error())
		}
		var ctx context.Context = nil
		_, err = client.GetDisplayNameWithContext(ctx, 35.77331417156089, 51.41831696033478, NewDefaultCallOptions(
			WithZoomLevel(17),
			WithFarsiLanguage(),
			WithPassengerResponseType(),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("there should be an error when creating request")
		}
	})
}
