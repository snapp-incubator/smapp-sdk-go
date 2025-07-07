package search

import (
	"context"
	_ "embed"
	"github.com/snapp-incubator/smapp-sdk-go/config"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

//go:embed search.json
var response []byte

func TestNewSearchClient(t *testing.T) {
	t.Run("without_options", func(t *testing.T) {
		cfg, err := config.NewDefaultConfig("key")
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewSearchClient(cfg, V1, time.Second)
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
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
		client, err := NewSearchClient(cfg, V1, time.Second, WithURL("https://google.com"))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		if client == nil {
			t.Fatalf("client should not be nil")
		}

		if client.url != "https://google.com" {
			t.Fatalf("client.url should be %s but it is %s", "https://google.com", client.url)
		}
	})
}

func TestClient_GetCitiesWithContext(t *testing.T) {
	t.Run("invalid_request", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{}`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.QueryParamSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewSearchClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		var ctx context.Context = nil
		_, err = client.GetCitiesWithContext(ctx, NewDefaultCallOptions())
		if err == nil {
			t.Fatalf("there should be an error when creating request")
		}
	})
}

func TestClient_SearchCityWithContext(t *testing.T) {
	t.Run("invalid_request", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{}`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.QueryParamSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewSearchClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		var ctx context.Context = nil
		_, err = client.SearchCityWithContext(ctx, "foo", NewDefaultCallOptions())
		if err == nil {
			t.Fatalf("there should be an error when creating request")
		}
	})
}

func TestClient_AutoCompleteWithContext(t *testing.T) {
	t.Run("invalid_request", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{}`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.QueryParamSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewSearchClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		var ctx context.Context = nil
		_, err = client.AutoCompleteWithContext(ctx, "foo", NewDefaultCallOptions())
		if err == nil {
			t.Fatalf("there should be an error when creating request")
		}
	})
}

func TestClient_DetailsWithContext(t *testing.T) {
	t.Run("invalid_request", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{}`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.QueryParamSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		client, err := NewSearchClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		var ctx context.Context = nil
		_, err = client.DetailsWithContext(ctx, "foo", NewDefaultCallOptions())
		if err == nil {
			t.Fatalf("there should be an error when creating request")
		}
	})
}

func TestClient_GetCities(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"predictions":[{"id":1000,"name":"تهران","centroid":{"latitude":"35.7006177","longitude":"51.4013785"},"description":"تهران"},{"id":1200,"name":"اصفهان","centroid":{"latitude":"32.6707877","longitude":"51.6650002"},"description":"اصفهان"},{"id":1300,"name":"شیراز","centroid":{"latitude":"29.6060218","longitude":"52.5378041"},"description":"فارس"},{"id":1400,"name":"مشهد","centroid":{"latitude":"36.2974945","longitude":"59.6059232"},"description":"خراسان رضوی"},{"id":1600,"name":"تبریز","centroid":{"latitude":"38.0739964","longitude":"46.2961952"},"description":"آذربایجان شرقی"}],"powered-by":"Smapp","status":"OK"}`))
		}))

		cfg, err := config.NewDefaultConfig("key")
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}

		client, err := NewSearchClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		cities, err := client.GetCities(NewDefaultCallOptions(
			WithFarsiLanguage(),
			WithOriginRequestContext(),
			WithLocation(50, 50),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err != nil {
			t.Fatalf("could not get cities: %s", err.Error())
		}
		if len(cities) != 5 {
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

		client, err := NewSearchClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		_, err = client.GetCities(NewDefaultCallOptions(
			WithFarsiLanguage(),
			WithOriginRequestContext(),
			WithLocation(50, 50),
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
			_, _ = w.Write([]byte(`{`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.QueryParamSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		cfg.APIKeySource = "foo"

		client, err := NewSearchClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		_, err = client.GetCities(NewDefaultCallOptions(
			WithFarsiLanguage(),
			WithOriginRequestContext(),
			WithLocation(50, 50),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("there should be an error when parsing request")
		}
	})
	t.Run("timeout", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(50 * time.Millisecond)
			_, _ = w.Write([]byte(`{"predictions":[{"id":1000,"name":"تهران","centroid":{"latitude":"35.7006177","longitude":"51.4013785"},"description":"تهران"},{"id":1200,"name":"اصفهان","centroid":{"latitude":"32.6707877","longitude":"51.6650002"},"description":"اصفهان"},{"id":1300,"name":"شیراز","centroid":{"latitude":"29.6060218","longitude":"52.5378041"},"description":"فارس"},{"id":1400,"name":"مشهد","centroid":{"latitude":"36.2974945","longitude":"59.6059232"},"description":"خراسان رضوی"},{"id":1600,"name":"تبریز","centroid":{"latitude":"38.0739964","longitude":"46.2961952"},"description":"آذربایجان شرقی"}],"powered-by":"Smapp","status":"OK"}`))
		}))

		cfg, err := config.NewDefaultConfig("key")
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}

		client, err := NewSearchClient(cfg, V1, 10*time.Millisecond, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		_, err = client.GetCities(NewDefaultCallOptions(
			WithFarsiLanguage(),
			WithOriginRequestContext(),
			WithLocation(50, 50),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("request should be timed out")
		}
	})
	t.Run("error_status", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"status": "ERROR"}`))
		}))

		cfg, err := config.NewDefaultConfig("key")
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}

		client, err := NewSearchClient(cfg, V1, 10*time.Millisecond, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		_, err = client.GetCities(NewDefaultCallOptions(
			WithFarsiLanguage(),
			WithOriginRequestContext(),
			WithLocation(50, 50),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("request status should not be ok")
		}
	})
	t.Run("non_200_status", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`{"predictions":[{"id":1000,"name":"تهران","centroid":{"latitude":"35.7006177","longitude":"51.4013785"},"description":"تهران"},{"id":1200,"name":"اصفهان","centroid":{"latitude":"32.6707877","longitude":"51.6650002"},"description":"اصفهان"},{"id":1300,"name":"شیراز","centroid":{"latitude":"29.6060218","longitude":"52.5378041"},"description":"فارس"},{"id":1400,"name":"مشهد","centroid":{"latitude":"36.2974945","longitude":"59.6059232"},"description":"خراسان رضوی"},{"id":1600,"name":"تبریز","centroid":{"latitude":"38.0739964","longitude":"46.2961952"},"description":"آذربایجان شرقی"}],"powered-by":"Smapp","status":"OK"}`))
		}))

		cfg, err := config.NewDefaultConfig("key")
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}

		client, err := NewSearchClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		_, err = client.GetCities(NewDefaultCallOptions(
			WithFarsiLanguage(),
			WithOriginRequestContext(),
			WithLocation(50, 50),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("request status should not be 200")
		}
	})
}

func TestClient_SearchCity(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"predictions":[{"id":1000,"name":"تهران","centroid":{"latitude":"35.7006177","longitude":"51.4013785"},"description":"تهران"},{"id":1200,"name":"اصفهان","centroid":{"latitude":"32.6707877","longitude":"51.6650002"},"description":"اصفهان"},{"id":1300,"name":"شیراز","centroid":{"latitude":"29.6060218","longitude":"52.5378041"},"description":"فارس"},{"id":1400,"name":"مشهد","centroid":{"latitude":"36.2974945","longitude":"59.6059232"},"description":"خراسان رضوی"},{"id":1600,"name":"تبریز","centroid":{"latitude":"38.0739964","longitude":"46.2961952"},"description":"آذربایجان شرقی"}],"powered-by":"Smapp","status":"OK"}`))
		}))

		cfg, err := config.NewDefaultConfig("key")
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}

		client, err := NewSearchClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		cities, err := client.SearchCity("تهران", NewDefaultCallOptions(
			WithFarsiLanguage(),
			WithOriginRequestContext(),
			WithLocation(50, 50),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err != nil {
			t.Fatalf("could not get cities: %s", err.Error())
		}
		if len(cities) != 5 {
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

		client, err := NewSearchClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		_, err = client.SearchCity("tehran", NewDefaultCallOptions(
			WithFarsiLanguage(),
			WithOriginRequestContext(),
			WithLocation(50, 50),
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
			_, _ = w.Write([]byte(`{`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.QueryParamSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		cfg.APIKeySource = "foo"

		client, err := NewSearchClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		_, err = client.SearchCity("tehran", NewDefaultCallOptions(
			WithFarsiLanguage(),
			WithOriginRequestContext(),
			WithLocation(50, 50),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("there should be an error when parsing request")
		}
	})
	t.Run("timeout", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(50 * time.Millisecond)
			_, _ = w.Write([]byte(`{"predictions":[{"id":1000,"name":"تهران","centroid":{"latitude":"35.7006177","longitude":"51.4013785"},"description":"تهران"},{"id":1200,"name":"اصفهان","centroid":{"latitude":"32.6707877","longitude":"51.6650002"},"description":"اصفهان"},{"id":1300,"name":"شیراز","centroid":{"latitude":"29.6060218","longitude":"52.5378041"},"description":"فارس"},{"id":1400,"name":"مشهد","centroid":{"latitude":"36.2974945","longitude":"59.6059232"},"description":"خراسان رضوی"},{"id":1600,"name":"تبریز","centroid":{"latitude":"38.0739964","longitude":"46.2961952"},"description":"آذربایجان شرقی"}],"powered-by":"Smapp","status":"OK"}`))
		}))

		cfg, err := config.NewDefaultConfig("key")
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}

		client, err := NewSearchClient(cfg, V1, 10*time.Millisecond, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		_, err = client.SearchCity("tehran", NewDefaultCallOptions(
			WithFarsiLanguage(),
			WithOriginRequestContext(),
			WithLocation(50, 50),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("request should be timed out")
		}
	})
	t.Run("error_status", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"status": "ERROR"}`))
		}))

		cfg, err := config.NewDefaultConfig("key")
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}

		client, err := NewSearchClient(cfg, V1, 10*time.Millisecond, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		_, err = client.SearchCity("tehran", NewDefaultCallOptions(
			WithFarsiLanguage(),
			WithOriginRequestContext(),
			WithLocation(50, 50),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("request status should not be ok")
		}
	})
	t.Run("non_200_status", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`{"predictions":[{"id":1000,"name":"تهران","centroid":{"latitude":"35.7006177","longitude":"51.4013785"},"description":"تهران"},{"id":1200,"name":"اصفهان","centroid":{"latitude":"32.6707877","longitude":"51.6650002"},"description":"اصفهان"},{"id":1300,"name":"شیراز","centroid":{"latitude":"29.6060218","longitude":"52.5378041"},"description":"فارس"},{"id":1400,"name":"مشهد","centroid":{"latitude":"36.2974945","longitude":"59.6059232"},"description":"خراسان رضوی"},{"id":1600,"name":"تبریز","centroid":{"latitude":"38.0739964","longitude":"46.2961952"},"description":"آذربایجان شرقی"}],"powered-by":"Smapp","status":"OK"}`))
		}))

		cfg, err := config.NewDefaultConfig("key")
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}

		client, err := NewSearchClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		_, err = client.SearchCity("tehran", NewDefaultCallOptions(
			WithFarsiLanguage(),
			WithOriginRequestContext(),
			WithLocation(50, 50),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("request status should not be 200")
		}
	})
}

func TestClient_AutoComplete(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write(response)
		}))

		cfg, err := config.NewDefaultConfig("key")
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}

		client, err := NewSearchClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		results, err := client.AutoComplete("آزادی", NewDefaultCallOptions(
			WithLocation(50, 50),
			WithUserLocation(60, 60),
			WithFarsiLanguage(),
			WithOriginRequestContext(),
			WithCityId(1000),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err != nil {
			t.Fatalf("could not get cities: %s", err.Error())
		}
		if len(results) == 0 {
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

		client, err := NewSearchClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		_, err = client.AutoComplete("tehran", NewDefaultCallOptions(
			WithFarsiLanguage(),
			WithOriginRequestContext(),
			WithLocation(50, 50),
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
			_, _ = w.Write([]byte(`{}`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.QueryParamSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		cfg.APIKeySource = "foo"

		client, err := NewSearchClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		_, err = client.AutoComplete("tehran", NewDefaultCallOptions(
			WithFarsiLanguage(),
			WithOriginRequestContext(),
			WithLocation(50, 50),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("there should be an error when parsing request")
		}
	})
	t.Run("timeout", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(50 * time.Millisecond)
			_, _ = w.Write(response)
		}))
		cfg, err := config.NewDefaultConfig("key")
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}

		client, err := NewSearchClient(cfg, V1, 10*time.Millisecond, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		_, err = client.AutoComplete("tehran", NewDefaultCallOptions(
			WithFarsiLanguage(),
			WithOriginRequestContext(),
			WithLocation(50, 50),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("request should be timed out")
		}
	})
	t.Run("error_status", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"status": "ERROR"}`))
		}))

		cfg, err := config.NewDefaultConfig("key")
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}

		client, err := NewSearchClient(cfg, V1, 10*time.Millisecond, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		_, err = client.AutoComplete("tehran", NewDefaultCallOptions(
			WithFarsiLanguage(),
			WithOriginRequestContext(),
			WithLocation(50, 50),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("request status should not be ok")
		}
	})
	t.Run("non_200_status", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`{"predictions":[{"id":1000,"name":"تهران","centroid":{"latitude":"35.7006177","longitude":"51.4013785"},"description":"تهران"},{"id":1200,"name":"اصفهان","centroid":{"latitude":"32.6707877","longitude":"51.6650002"},"description":"اصفهان"},{"id":1300,"name":"شیراز","centroid":{"latitude":"29.6060218","longitude":"52.5378041"},"description":"فارس"},{"id":1400,"name":"مشهد","centroid":{"latitude":"36.2974945","longitude":"59.6059232"},"description":"خراسان رضوی"},{"id":1600,"name":"تبریز","centroid":{"latitude":"38.0739964","longitude":"46.2961952"},"description":"آذربایجان شرقی"}],"powered-by":"Smapp","status":"OK"}`))
		}))

		cfg, err := config.NewDefaultConfig("key")
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}

		client, err := NewSearchClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		_, err = client.AutoComplete("tehran", NewDefaultCallOptions(
			WithFarsiLanguage(),
			WithOriginRequestContext(),
			WithLocation(50, 50),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("request status should not be 200")
		}
	})
}

func TestClient_Details(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"status":"OK","result":{"name":"میدان آزادی - سبزه میدان","geometry":{"location":{"lat":36.2694934,"lng":50.0041848}}}}`))
		}))

		cfg, err := config.NewDefaultConfig("key")
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}

		client, err := NewSearchClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		detail, err := client.Details("<string>::36491302070", NewDefaultCallOptions(
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err != nil {
			t.Fatalf("could not get cities: %s", err.Error())
		}
		if detail.Name == "" {
			t.Fatalf("detail name should not be empty")
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

		client, err := NewSearchClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		_, err = client.Details("id", NewDefaultCallOptions(
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
			_, _ = w.Write([]byte(`{}`))
		}))

		cfg, err := config.NewDefaultConfig("key", config.WithAPIKeySource(config.QueryParamSource))
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}
		cfg.APIKeySource = "foo"

		client, err := NewSearchClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		_, err = client.Details("tehran", NewDefaultCallOptions(
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("there should be an error when parsing request")
		}
	})
	t.Run("timeout", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(50 * time.Millisecond)
			_, _ = w.Write([]byte(`{"status":"OK","result":{"name":"میدان آزادی - سبزه میدان","geometry":{"location":{"lat":36.2694934,"lng":50.0041848}}}}`))
		}))

		cfg, err := config.NewDefaultConfig("key")
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}

		client, err := NewSearchClient(cfg, V1, 10*time.Millisecond, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		_, err = client.Details("tehran", NewDefaultCallOptions(
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("request should be timed out")
		}
	})
	t.Run("non_200_status", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`{"status":"OK","result":{"name":"میدان آزادی - سبزه میدان","geometry":{"location":{"lat":36.2694934,"lng":50.0041848}}}}`))
		}))

		cfg, err := config.NewDefaultConfig("key")
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}

		client, err := NewSearchClient(cfg, V1, time.Second, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		_, err = client.Details("tehran", NewDefaultCallOptions(
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("request status should not be 200")
		}
	})
}

func TestClient_Details_ErrorStatus(t *testing.T) {
	t.Run("error_status", func(t *testing.T) {
		sv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"status": "ERROR"}`))
		}))

		cfg, err := config.NewDefaultConfig("key")
		if err != nil {
			t.Fatalf("could not create default config due to: %s", err.Error())
		}

		client, err := NewSearchClient(cfg, V1, 10*time.Millisecond, WithURL(sv.URL))
		if err != nil {
			t.Fatalf("could not create search client due to: %s", err.Error())
		}
		_, err = client.Details("id", NewDefaultCallOptions(
			WithLocation(50, 50),
			WithHeaders(map[string]string{
				"foo": "bar",
			}),
		))
		if err == nil {
			t.Fatalf("request status should not be ok")
		}
	})
}
