package reverse

import "testing"

func TestWithDriverResponseType(t *testing.T) {
	callOptions := CallOptions{
		ZoomLevel:    16,
		ResponseType: Passenger,
		Language:     Farsi,
	}

	WithDriverResponseType()(&callOptions)

	if callOptions.ResponseType != Driver {
		t.Fatalf("ResponseType should be Driver")
	}
}

func TestWithPassengerResponseType(t *testing.T) {
	callOptions := CallOptions{
		ZoomLevel:    16,
		ResponseType: Driver,
		Language:     Farsi,
	}

	WithPassengerResponseType()(&callOptions)

	if callOptions.ResponseType != Passenger {
		t.Fatalf("ResponseType should be Passenger")
	}
}

func TestWithVerboseResponseType(t *testing.T) {
	callOptions := CallOptions{
		ZoomLevel:    16,
		ResponseType: Driver,
		Language:     Farsi,
	}

	WithVerboseResponseType()(&callOptions)

	if callOptions.ResponseType != Verbose {
		t.Fatalf("ResponseType should be Verbose")
	}
}

func TestWithOriginResponseType(t *testing.T) {
	callOptions := CallOptions{
		ZoomLevel:    16,
		ResponseType: Driver,
		Language:     Farsi,
	}

	WithOriginResponseType()(&callOptions)

	if callOptions.ResponseType != Origin {
		t.Fatalf("ResponseType should be Origin")
	}
}

func TestWithDestinationResponseType(t *testing.T) {
	callOptions := CallOptions{
		ZoomLevel:    16,
		ResponseType: Driver,
		Language:     Farsi,
	}

	WithDestinationResponseType()(&callOptions)

	if callOptions.ResponseType != Destination {
		t.Fatalf("ResponseType should be Destination")
	}
}

func TestWithIraqResponseType(t *testing.T) {
	callOptions := CallOptions{
		ZoomLevel:    16,
		ResponseType: Driver,
		Language:     Farsi,
	}

	WithIraqResponseType()(&callOptions)

	if callOptions.ResponseType != Iraq {
		t.Fatalf("ResponseType should be Iraq")
	}
}

func TestWithBikerResponseType(t *testing.T) {
	callOptions := CallOptions{
		ZoomLevel:    16,
		ResponseType: Driver,
		Language:     Farsi,
	}

	WithBikerResponseType()(&callOptions)

	if callOptions.ResponseType != Biker {
		t.Fatalf("ResponseType should be Biker")
	}
}

func TestWithFarsiLanguage(t *testing.T) {
	callOptions := CallOptions{
		ZoomLevel:    16,
		ResponseType: Driver,
		Language:     English,
	}

	WithFarsiLanguage()(&callOptions)

	if callOptions.Language != Farsi {
		t.Fatalf("Language should be Farsi")
	}
}

func TestWithEnglishLanguage(t *testing.T) {
	callOptions := CallOptions{
		ZoomLevel:    16,
		ResponseType: Driver,
		Language:     Farsi,
	}

	WithEnglishLanguage()(&callOptions)

	if callOptions.Language != English {
		t.Fatalf("Language should be English")
	}
}

func TestWithArabicLanguage(t *testing.T) {
	callOptions := CallOptions{
		ZoomLevel:    16,
		ResponseType: Driver,
		Language:     Farsi,
	}

	WithArabicLanguage()(&callOptions)

	if callOptions.Language != Arabic {
		t.Fatalf("Language should be Arabic")
	}
}

func TestWithZoomLevel(t *testing.T) {
	callOptions := CallOptions{
		ZoomLevel:    16,
		ResponseType: Driver,
		Language:     Farsi,
	}

	WithZoomLevel(10)(&callOptions)

	if callOptions.ZoomLevel != 10 {
		t.Fatalf("ZoomLevel should be 10")
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

func TestNewDefaultCallOptions(t *testing.T) {
	callOptions := NewDefaultCallOptions(WithZoomLevel(10))

	if callOptions.ZoomLevel != 10 {
		t.Fatalf("ZoomLevel should be 10")
	}
}
