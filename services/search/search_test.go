package search

import (
	"context"
	"gitlab.snapp.ir/Map/sdk/smapp-sdk-go/config"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

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
			_, _ = w.Write([]byte(`{"predictions":[{"place_id":"<string>::36491302070","name":"میدان آزادی - سبزه میدان (سبزه میدان)","description":"قزوین، بلاغی، میدان آزادی","structured_formatting":{"main_text":"میدان آزادی - سبزه میدان (سبزه میدان)","secondary_text":"قزوین، بلاغی، میدان آزادی"},"type":"place:locality","location":{"latitude":"36.2694934","longitude":"50.0041848"},"distance":6536672,"area_length":0,"all_tags":["place:locality"]},{"place_id":"<string>::950710101","name":"میدان آزادی","description":"درگز، نزدیک آزادی","structured_formatting":{"main_text":"میدان آزادی","secondary_text":"درگز، نزدیک آزادی"},"type":"junction:roundabout","location":{"latitude":"37.442597649999996","longitude":"59.107219395636754"},"distance":7332163,"area_length":994.9650155398995,"all_tags":["highway:primary","junction:roundabout"]},{"place_id":"<string>::9140361451","name":"میدان آزادی (دروازه شیراز)","description":"اصفهان، بهار آزادی، نزدیک بهار آزادی","structured_formatting":{"main_text":"میدان آزادی (دروازه شیراز)","secondary_text":"اصفهان، بهار آزادی، نزدیک بهار آزادی"},"type":"junction:roundabout","location":{"latitude":"32.622317249999995","longitude":"51.66451498034586"},"distance":6505225,"area_length":16798.216844710798,"all_tags":["highway:primary","junction:roundabout"]},{"place_id":"<string>::4645831551","name":"آزادی","description":"تهران، منطقه ۱۰، نزدیک کنارگذر آزادی","structured_formatting":{"main_text":"آزادی","secondary_text":"تهران، منطقه ۱۰، نزدیک کنارگذر آزادی"},"type":"highway:primary","location":{"latitude":"35.7002503","longitude":"51.3636759"},"distance":6619441,"area_length":616.3962327791401,"all_tags":["highway:primary"]},{"place_id":"<string>::4551093451","name":"آزادی","description":"تهران، منطقه ۹، نزدیک کنارگذر آزادی","structured_formatting":{"main_text":"آزادی","secondary_text":"تهران، منطقه ۹، نزدیک کنارگذر آزادی"},"type":"highway:primary","location":{"latitude":"35.6998263","longitude":"51.3452595"},"distance":6617914,"area_length":1167.3216585595562,"all_tags":["highway:primary"]},{"place_id":"<string>::8842185051","name":"آزادی","description":"تهران، منطقه ۲، تیموری، نزدیک کنارگذر آزادی","structured_formatting":{"main_text":"آزادی","secondary_text":"تهران، منطقه ۲، تیموری، نزدیک کنارگذر آزادی"},"type":"highway:primary","location":{"latitude":"35.700128","longitude":"51.3522168"},"distance":6618498,"area_length":37.66038676593313,"all_tags":["highway:primary"]},{"place_id":"<string>::257769021","name":"آزادی","description":"تهران، نزدیک کنارگذر آزادی","structured_formatting":{"main_text":"آزادی","secondary_text":"تهران، نزدیک کنارگذر آزادی"},"type":"highway:primary","location":{"latitude":"35.6999819","longitude":"51.3454527"},"distance":6617937,"area_length":1154.204709552337,"all_tags":["highway:primary"]},{"place_id":"<string>::8842168051","name":"آزادی","description":"تهران، منطقه ۲، زنجان شمالی، نزدیک کنارگذر آزادی","structured_formatting":{"main_text":"آزادی","secondary_text":"تهران، منطقه ۲، زنجان شمالی، نزدیک کنارگذر آزادی"},"type":"highway:primary","location":{"latitude":"35.7002159","longitude":"51.3561173"},"distance":6618821,"area_length":352.0953999942009,"all_tags":["highway:primary"]},{"place_id":"<string>::7124789821","name":"آزادی","description":"تهران، منطقه ۲، شادمان، نزدیک کنارگذر آزادی","structured_formatting":{"main_text":"آزادی","secondary_text":"تهران، منطقه ۲، شادمان، نزدیک کنارگذر آزادی"},"type":"highway:primary","location":{"latitude":"35.7003092","longitude":"51.3603219"},"distance":6619170,"area_length":237.41218608846853,"all_tags":["highway:primary"]},{"place_id":"<string>::8842168021","name":"آزادی","description":"تهران، منطقه ۱۰، زنجان جنوبی، نزدیک کنارگذر آزادی","structured_formatting":{"main_text":"آزادی","secondary_text":"تهران، منطقه ۱۰، زنجان جنوبی، نزدیک کنارگذر آزادی"},"type":"highway:primary","location":{"latitude":"35.7000742","longitude":"51.3555612"},"distance":6618769,"area_length":23.36971995455747,"all_tags":["highway:primary"]},{"place_id":"<string>::257769011","name":"آزادی","description":"تهران، منطقه ۲، نزدیک کنارگذر آزادی","structured_formatting":{"main_text":"آزادی","secondary_text":"تهران، منطقه ۲، نزدیک کنارگذر آزادی"},"type":"highway:primary","location":{"latitude":"35.7003092","longitude":"51.3603219"},"distance":6619170,"area_length":78.18902705212697,"all_tags":["tunnel:yes","highway:primary"]},{"place_id":"<string>::4831879351","name":"آزادی","description":"ساوه، عبدل آباد شرقی، نزدیک ۲۴ آزادی","structured_formatting":{"main_text":"آزادی","secondary_text":"ساوه، عبدل آباد شرقی، نزدیک ۲۴ آزادی"},"type":"highway:primary","location":{"latitude":"35.0416674","longitude":"50.3704184"},"distance":6506770,"area_length":1615.3377548218934,"all_tags":["highway:primary"]},{"place_id":"<string>::8842185041","name":"آزادی","description":"تهران، منطقه ۹، دکتر هوشیار، نزدیک کنارگذر آزادی","structured_formatting":{"main_text":"آزادی","secondary_text":"تهران، منطقه ۹، دکتر هوشیار، نزدیک کنارگذر آزادی"},"type":"highway:primary","location":{"latitude":"35.6999794","longitude":"51.35265"},"distance":6618526,"area_length":30.545487935962196,"all_tags":["highway:primary"]},{"place_id":"<string>::3672419121","name":"بلوار آزادی","description":"ارومیه، مجسمه، نزدیک آزادی","structured_formatting":{"main_text":"بلوار آزادی","secondary_text":"ارومیه، مجسمه، نزدیک آزادی"},"type":"highway:primary","location":{"latitude":"37.5776308","longitude":"45.0621208"},"distance":6221997,"area_length":1178.5079588589217,"all_tags":["highway:primary"]},{"place_id":"<string>::4982625351","name":"بلوار آزادی","description":"سوادکوه، نزدیک آزادی ۲۰","structured_formatting":{"main_text":"بلوار آزادی","secondary_text":"سوادکوه، نزدیک آزادی ۲۰"},"type":"highway:trunk","location":{"latitude":"36.1687295","longitude":"52.9952702"},"distance":6774706,"area_length":1526.7825736316615,"all_tags":["highway:trunk"]},{"place_id":"<string>::3826297861","name":"بزرگراه آزادی","description":"مشهد، منطقه ۱۰، خاتم الانبیا، نزدیک آزادی ۱۰۳","structured_formatting":{"main_text":"بزرگراه آزادی","secondary_text":"مشهد، منطقه ۱۰، خاتم الانبیا، نزدیک آزادی ۱۰۳"},"type":"highway:trunk","location":{"latitude":"36.371039","longitude":"59.5330252"},"distance":7327838,"area_length":50.08785146468531,"all_tags":["highway:trunk"]},{"place_id":"<string>::3818126361","name":"بزرگراه آزادی","description":"مشهد، منطقه ۱۱، آزادشهر، نزدیک آزادی ۹","structured_formatting":{"main_text":"بزرگراه آزادی","secondary_text":"مشهد، منطقه ۱۱، آزادشهر، نزدیک آزادی ۹"},"type":"highway:trunk","location":{"latitude":"36.3175705","longitude":"59.5411516"},"distance":7326568,"area_length":106.75699513632644,"all_tags":["highway:trunk"]},{"place_id":"<string>::4940116131","name":"بزرگراه آزادی","description":"مشهد، منطقه ۱۱، آزادشهر، نزدیک کوچه آزادی ۲۴","structured_formatting":{"main_text":"بزرگراه آزادی","secondary_text":"مشهد، منطقه ۱۱، آزادشهر، نزدیک کوچه آزادی ۲۴"},"type":"highway:trunk","location":{"latitude":"36.3315627","longitude":"59.5447683"},"distance":7327386,"area_length":641.0389255255227,"all_tags":["highway:trunk"]},{"place_id":"<string>::9500940301","name":"بزرگراه آزادی","description":"مشهد، منطقه ۱۱، آزادشهر، نزدیک آزادی ۲۳","structured_formatting":{"main_text":"بزرگراه آزادی","secondary_text":"مشهد، منطقه ۱۱، آزادشهر، نزدیک آزادی ۲۳"},"type":"highway:trunk","location":{"latitude":"36.3231456","longitude":"59.5425723"},"distance":7326892,"area_length":76.90050667479204,"all_tags":["highway:trunk"]},{"place_id":"<string>::4940116161","name":"بزرگراه آزادی","description":"مشهد، منطقه ۲، شهید فرامرز عباسی، نزدیک ۵۳ آزادی","structured_formatting":{"main_text":"بزرگراه آزادی","secondary_text":"مشهد، منطقه ۲، شهید فرامرز عباسی، نزدیک ۵۳ آزادی"},"type":"highway:trunk","location":{"latitude":"36.3315312","longitude":"59.5450254"},"distance":7327406,"area_length":420.3062252980578,"all_tags":["highway:trunk"]},{"place_id":"<string>::4895019581","name":"میدان آزادی (انوش)","description":"خرم‌آباد، آزادی، نزدیک بلوار شریعتی","structured_formatting":{"main_text":"میدان آزادی (انوش)","secondary_text":"خرم‌آباد، آزادی، نزدیک بلوار شریعتی"},"type":"junction:roundabout","location":{"latitude":"33.486051149999994","longitude":"48.35839134257917"},"distance":6265054,"area_length":682.5737797801849,"all_tags":["highway:secondary","junction:roundabout"]},{"place_id":"<string>::6972644111","name":"آزادی","description":"قائمشهر، شماره ۲، نزدیک آزادی ۱۰۱","structured_formatting":{"main_text":"آزادی","secondary_text":"قائمشهر، شماره ۲، نزدیک آزادی ۱۰۱"},"type":"highway:tertiary","location":{"latitude":"36.4807771","longitude":"52.8969535"},"distance":6780832,"area_length":83.37227712524648,"all_tags":["highway:tertiary"]},{"place_id":"<string>::1246426231","name":"آزادی","description":"زاهدان، نزدیک آزادی ۱۲","structured_formatting":{"main_text":"آزادی","secondary_text":"زاهدان، نزدیک آزادی ۱۲"},"type":"highway:secondary","location":{"latitude":"29.4993277","longitude":"60.8717227"},"distance":7219977,"area_length":530.118310140061,"all_tags":["highway:secondary"]},{"place_id":"<string>::6658797651","name":"آزادی","description":"قائمشهر، دخانیات، نزدیک آزادی ۵۹ شهید یوسفی","structured_formatting":{"main_text":"آزادی","secondary_text":"قائمشهر، دخانیات، نزدیک آزادی ۵۹ شهید یوسفی"},"type":"highway:tertiary","location":{"latitude":"36.4532419","longitude":"52.8517442"},"distance":6775889,"area_length":1798.4848183711292,"all_tags":["highway:tertiary"]},{"place_id":"<string>::2432461751","name":"آزادی","description":"قائمشهر، نزدیک آزادی ۷۳","structured_formatting":{"main_text":"آزادی","secondary_text":"قائمشهر، نزدیک آزادی ۷۳"},"type":"highway:tertiary","location":{"latitude":"36.4663208","longitude":"52.86981"},"distance":6777959,"area_length":137.228571169009,"all_tags":["highway:tertiary"]}],"powered-by":"Smapp","status":"OK"}`))
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
			_, _ = w.Write([]byte(`{"predictions":[{"place_id":"<string>::36491302070","name":"میدان آزادی - سبزه میدان (سبزه میدان)","description":"قزوین، بلاغی، میدان آزادی","structured_formatting":{"main_text":"میدان آزادی - سبزه میدان (سبزه میدان)","secondary_text":"قزوین، بلاغی، میدان آزادی"},"type":"place:locality","location":{"latitude":"36.2694934","longitude":"50.0041848"},"distance":6536672,"area_length":0,"all_tags":["place:locality"]},{"place_id":"<string>::950710101","name":"میدان آزادی","description":"درگز، نزدیک آزادی","structured_formatting":{"main_text":"میدان آزادی","secondary_text":"درگز، نزدیک آزادی"},"type":"junction:roundabout","location":{"latitude":"37.442597649999996","longitude":"59.107219395636754"},"distance":7332163,"area_length":994.9650155398995,"all_tags":["highway:primary","junction:roundabout"]},{"place_id":"<string>::9140361451","name":"میدان آزادی (دروازه شیراز)","description":"اصفهان، بهار آزادی، نزدیک بهار آزادی","structured_formatting":{"main_text":"میدان آزادی (دروازه شیراز)","secondary_text":"اصفهان، بهار آزادی، نزدیک بهار آزادی"},"type":"junction:roundabout","location":{"latitude":"32.622317249999995","longitude":"51.66451498034586"},"distance":6505225,"area_length":16798.216844710798,"all_tags":["highway:primary","junction:roundabout"]},{"place_id":"<string>::4645831551","name":"آزادی","description":"تهران، منطقه ۱۰، نزدیک کنارگذر آزادی","structured_formatting":{"main_text":"آزادی","secondary_text":"تهران، منطقه ۱۰، نزدیک کنارگذر آزادی"},"type":"highway:primary","location":{"latitude":"35.7002503","longitude":"51.3636759"},"distance":6619441,"area_length":616.3962327791401,"all_tags":["highway:primary"]},{"place_id":"<string>::4551093451","name":"آزادی","description":"تهران، منطقه ۹، نزدیک کنارگذر آزادی","structured_formatting":{"main_text":"آزادی","secondary_text":"تهران، منطقه ۹، نزدیک کنارگذر آزادی"},"type":"highway:primary","location":{"latitude":"35.6998263","longitude":"51.3452595"},"distance":6617914,"area_length":1167.3216585595562,"all_tags":["highway:primary"]},{"place_id":"<string>::8842185051","name":"آزادی","description":"تهران، منطقه ۲، تیموری، نزدیک کنارگذر آزادی","structured_formatting":{"main_text":"آزادی","secondary_text":"تهران، منطقه ۲، تیموری، نزدیک کنارگذر آزادی"},"type":"highway:primary","location":{"latitude":"35.700128","longitude":"51.3522168"},"distance":6618498,"area_length":37.66038676593313,"all_tags":["highway:primary"]},{"place_id":"<string>::257769021","name":"آزادی","description":"تهران، نزدیک کنارگذر آزادی","structured_formatting":{"main_text":"آزادی","secondary_text":"تهران، نزدیک کنارگذر آزادی"},"type":"highway:primary","location":{"latitude":"35.6999819","longitude":"51.3454527"},"distance":6617937,"area_length":1154.204709552337,"all_tags":["highway:primary"]},{"place_id":"<string>::8842168051","name":"آزادی","description":"تهران، منطقه ۲، زنجان شمالی، نزدیک کنارگذر آزادی","structured_formatting":{"main_text":"آزادی","secondary_text":"تهران، منطقه ۲، زنجان شمالی، نزدیک کنارگذر آزادی"},"type":"highway:primary","location":{"latitude":"35.7002159","longitude":"51.3561173"},"distance":6618821,"area_length":352.0953999942009,"all_tags":["highway:primary"]},{"place_id":"<string>::7124789821","name":"آزادی","description":"تهران، منطقه ۲، شادمان، نزدیک کنارگذر آزادی","structured_formatting":{"main_text":"آزادی","secondary_text":"تهران، منطقه ۲، شادمان، نزدیک کنارگذر آزادی"},"type":"highway:primary","location":{"latitude":"35.7003092","longitude":"51.3603219"},"distance":6619170,"area_length":237.41218608846853,"all_tags":["highway:primary"]},{"place_id":"<string>::8842168021","name":"آزادی","description":"تهران، منطقه ۱۰، زنجان جنوبی، نزدیک کنارگذر آزادی","structured_formatting":{"main_text":"آزادی","secondary_text":"تهران، منطقه ۱۰، زنجان جنوبی، نزدیک کنارگذر آزادی"},"type":"highway:primary","location":{"latitude":"35.7000742","longitude":"51.3555612"},"distance":6618769,"area_length":23.36971995455747,"all_tags":["highway:primary"]},{"place_id":"<string>::257769011","name":"آزادی","description":"تهران، منطقه ۲، نزدیک کنارگذر آزادی","structured_formatting":{"main_text":"آزادی","secondary_text":"تهران، منطقه ۲، نزدیک کنارگذر آزادی"},"type":"highway:primary","location":{"latitude":"35.7003092","longitude":"51.3603219"},"distance":6619170,"area_length":78.18902705212697,"all_tags":["tunnel:yes","highway:primary"]},{"place_id":"<string>::4831879351","name":"آزادی","description":"ساوه، عبدل آباد شرقی، نزدیک ۲۴ آزادی","structured_formatting":{"main_text":"آزادی","secondary_text":"ساوه، عبدل آباد شرقی، نزدیک ۲۴ آزادی"},"type":"highway:primary","location":{"latitude":"35.0416674","longitude":"50.3704184"},"distance":6506770,"area_length":1615.3377548218934,"all_tags":["highway:primary"]},{"place_id":"<string>::8842185041","name":"آزادی","description":"تهران، منطقه ۹، دکتر هوشیار، نزدیک کنارگذر آزادی","structured_formatting":{"main_text":"آزادی","secondary_text":"تهران، منطقه ۹، دکتر هوشیار، نزدیک کنارگذر آزادی"},"type":"highway:primary","location":{"latitude":"35.6999794","longitude":"51.35265"},"distance":6618526,"area_length":30.545487935962196,"all_tags":["highway:primary"]},{"place_id":"<string>::3672419121","name":"بلوار آزادی","description":"ارومیه، مجسمه، نزدیک آزادی","structured_formatting":{"main_text":"بلوار آزادی","secondary_text":"ارومیه، مجسمه، نزدیک آزادی"},"type":"highway:primary","location":{"latitude":"37.5776308","longitude":"45.0621208"},"distance":6221997,"area_length":1178.5079588589217,"all_tags":["highway:primary"]},{"place_id":"<string>::4982625351","name":"بلوار آزادی","description":"سوادکوه، نزدیک آزادی ۲۰","structured_formatting":{"main_text":"بلوار آزادی","secondary_text":"سوادکوه، نزدیک آزادی ۲۰"},"type":"highway:trunk","location":{"latitude":"36.1687295","longitude":"52.9952702"},"distance":6774706,"area_length":1526.7825736316615,"all_tags":["highway:trunk"]},{"place_id":"<string>::3826297861","name":"بزرگراه آزادی","description":"مشهد، منطقه ۱۰، خاتم الانبیا، نزدیک آزادی ۱۰۳","structured_formatting":{"main_text":"بزرگراه آزادی","secondary_text":"مشهد، منطقه ۱۰، خاتم الانبیا، نزدیک آزادی ۱۰۳"},"type":"highway:trunk","location":{"latitude":"36.371039","longitude":"59.5330252"},"distance":7327838,"area_length":50.08785146468531,"all_tags":["highway:trunk"]},{"place_id":"<string>::3818126361","name":"بزرگراه آزادی","description":"مشهد، منطقه ۱۱، آزادشهر، نزدیک آزادی ۹","structured_formatting":{"main_text":"بزرگراه آزادی","secondary_text":"مشهد، منطقه ۱۱، آزادشهر، نزدیک آزادی ۹"},"type":"highway:trunk","location":{"latitude":"36.3175705","longitude":"59.5411516"},"distance":7326568,"area_length":106.75699513632644,"all_tags":["highway:trunk"]},{"place_id":"<string>::4940116131","name":"بزرگراه آزادی","description":"مشهد، منطقه ۱۱، آزادشهر، نزدیک کوچه آزادی ۲۴","structured_formatting":{"main_text":"بزرگراه آزادی","secondary_text":"مشهد، منطقه ۱۱، آزادشهر، نزدیک کوچه آزادی ۲۴"},"type":"highway:trunk","location":{"latitude":"36.3315627","longitude":"59.5447683"},"distance":7327386,"area_length":641.0389255255227,"all_tags":["highway:trunk"]},{"place_id":"<string>::9500940301","name":"بزرگراه آزادی","description":"مشهد، منطقه ۱۱، آزادشهر، نزدیک آزادی ۲۳","structured_formatting":{"main_text":"بزرگراه آزادی","secondary_text":"مشهد، منطقه ۱۱، آزادشهر، نزدیک آزادی ۲۳"},"type":"highway:trunk","location":{"latitude":"36.3231456","longitude":"59.5425723"},"distance":7326892,"area_length":76.90050667479204,"all_tags":["highway:trunk"]},{"place_id":"<string>::4940116161","name":"بزرگراه آزادی","description":"مشهد، منطقه ۲، شهید فرامرز عباسی، نزدیک ۵۳ آزادی","structured_formatting":{"main_text":"بزرگراه آزادی","secondary_text":"مشهد، منطقه ۲، شهید فرامرز عباسی، نزدیک ۵۳ آزادی"},"type":"highway:trunk","location":{"latitude":"36.3315312","longitude":"59.5450254"},"distance":7327406,"area_length":420.3062252980578,"all_tags":["highway:trunk"]},{"place_id":"<string>::4895019581","name":"میدان آزادی (انوش)","description":"خرم‌آباد، آزادی، نزدیک بلوار شریعتی","structured_formatting":{"main_text":"میدان آزادی (انوش)","secondary_text":"خرم‌آباد، آزادی، نزدیک بلوار شریعتی"},"type":"junction:roundabout","location":{"latitude":"33.486051149999994","longitude":"48.35839134257917"},"distance":6265054,"area_length":682.5737797801849,"all_tags":["highway:secondary","junction:roundabout"]},{"place_id":"<string>::6972644111","name":"آزادی","description":"قائمشهر، شماره ۲، نزدیک آزادی ۱۰۱","structured_formatting":{"main_text":"آزادی","secondary_text":"قائمشهر، شماره ۲، نزدیک آزادی ۱۰۱"},"type":"highway:tertiary","location":{"latitude":"36.4807771","longitude":"52.8969535"},"distance":6780832,"area_length":83.37227712524648,"all_tags":["highway:tertiary"]},{"place_id":"<string>::1246426231","name":"آزادی","description":"زاهدان، نزدیک آزادی ۱۲","structured_formatting":{"main_text":"آزادی","secondary_text":"زاهدان، نزدیک آزادی ۱۲"},"type":"highway:secondary","location":{"latitude":"29.4993277","longitude":"60.8717227"},"distance":7219977,"area_length":530.118310140061,"all_tags":["highway:secondary"]},{"place_id":"<string>::6658797651","name":"آزادی","description":"قائمشهر، دخانیات، نزدیک آزادی ۵۹ شهید یوسفی","structured_formatting":{"main_text":"آزادی","secondary_text":"قائمشهر، دخانیات، نزدیک آزادی ۵۹ شهید یوسفی"},"type":"highway:tertiary","location":{"latitude":"36.4532419","longitude":"52.8517442"},"distance":6775889,"area_length":1798.4848183711292,"all_tags":["highway:tertiary"]},{"place_id":"<string>::2432461751","name":"آزادی","description":"قائمشهر، نزدیک آزادی ۷۳","structured_formatting":{"main_text":"آزادی","secondary_text":"قائمشهر، نزدیک آزادی ۷۳"},"type":"highway:tertiary","location":{"latitude":"36.4663208","longitude":"52.86981"},"distance":6777959,"area_length":137.228571169009,"all_tags":["highway:tertiary"]}],"powered-by":"Smapp","status":"OK"}`))
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
