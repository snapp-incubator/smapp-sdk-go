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
