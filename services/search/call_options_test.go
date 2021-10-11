package search

import "testing"

func TestWithFarsiLanguage(t *testing.T) {
	callOptions := CallOptions{
		Language: English,
	}

	WithFarsiLanguage()(&callOptions)

	if callOptions.Language != Farsi {
		t.Fatalf("Language should be Farsi")
	}
}

func TestWithEnglishLanguage(t *testing.T) {
	callOptions := CallOptions{
		Language: Farsi,
	}

	WithEnglishLanguage()(&callOptions)

	if callOptions.Language != English {
		t.Fatalf("Language should be English")
	}
}

func TestWithLocation(t *testing.T) {
	callOptions := CallOptions{}

	WithLocation(50, 60)(&callOptions)

	if callOptions.UseLocation == false {
		t.Fatalf("Uselocation should be true")
	}

	if callOptions.Location.Lat != 50 || callOptions.Location.Lon != 60 {
		t.Fatalf("Invalid Location")
	}
}

func TestWithUserLocation(t *testing.T) {
	callOptions := CallOptions{}

	WithUserLocation(50, 60)(&callOptions)

	if callOptions.UseUserLocation == false {
		t.Fatalf("UseUserlocation should be true")
	}

	if callOptions.UserLocation.Lat != 50 || callOptions.UserLocation.Lon != 60 {
		t.Fatalf("Invalid UserLocation")
	}
}

func TestWithCityId(t *testing.T) {
	callOptions := CallOptions{}

	WithCityId(1000)(&callOptions)

	if callOptions.UseCityID == false {
		t.Fatalf("UseCityID should be true")
	}

	if callOptions.CityID != 1000 {
		t.Fatalf("Invalid CityID")
	}
}

func TestWithHeaders(t *testing.T) {
	callOptions := CallOptions{}

	WithHeaders(map[string]string{
		"foo": "bar",
	})(&callOptions)

	if callOptions.Headers == nil {
		t.Fatalf("Headers should not be nil")
	}

	if callOptions.Headers["foo"] != "bar" {
		t.Fatalf("Invalid Headers")
	}
}

func TestWithOriginRequestContext(t *testing.T) {
	callOptions := CallOptions{}

	WithOriginRequestContext()(&callOptions)

	if callOptions.RequestContext != Origin {
		t.Fatalf("RequestContext should be Origin")
	}
}

func TestWithFavouriteRequestContext(t *testing.T) {
	callOptions := CallOptions{}

	WithFavouriteRequestContext()(&callOptions)

	if callOptions.RequestContext != Favourite {
		t.Fatalf("RequestContext should be Favourite")
	}
}

func TestWithFirstDestinationRequestContext(t *testing.T) {
	callOptions := CallOptions{}

	WithFirstDestinationRequestContext()(&callOptions)

	if callOptions.RequestContext != FirstDestination {
		t.Fatalf("RequestContext should be FirstDestination")
	}
}

func TestWithSecondDestinationRequestContext(t *testing.T) {
	callOptions := CallOptions{}

	WithSecondDestinationRequestContext()(&callOptions)

	if callOptions.RequestContext != SecondDestination {
		t.Fatalf("RequestContext should be SecondDestination")
	}
}

func TestNewDefaultCallOptions(t *testing.T) {
	callOptions := NewDefaultCallOptions(WithCityId(1000))

	if callOptions.UseCityID == false {
		t.Fatalf("UseCityID should be true")
	}

	if callOptions.CityID != 1000 {
		t.Fatalf("Invalid CityID")
	}
}
