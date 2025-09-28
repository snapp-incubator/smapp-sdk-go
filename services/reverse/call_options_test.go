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

func TestWithDriverDestinationResponseType(t *testing.T) {
	callOptions := CallOptions{
		ZoomLevel:    16,
		ResponseType: Driver_Destination,
		Language:     Farsi,
	}

	WithDriverDestinationResponseType()(&callOptions)

	if callOptions.ResponseType != Driver_Destination {
		t.Fatalf("ResponseType should be Driver Destination")
	}
}

func TestWithDriverOriginResponseType(t *testing.T) {
	callOptions := CallOptions{
		ZoomLevel:    16,
		ResponseType: Driver_Origin,
		Language:     Farsi,
	}

	WithDriverOriginResponseType()(&callOptions)

	if callOptions.ResponseType != Driver_Origin {
		t.Fatalf("ResponseType should be Driver Origin")
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
func TestWithKurdishLanguage(t *testing.T) {
	callOptions := CallOptions{
		ZoomLevel:    16,
		ResponseType: Driver,
		Language:     Farsi,
	}

	WithKurdishLanguage()(&callOptions)

	if callOptions.Language != Kurdish {
		t.Fatalf("Language should be Kurdish")
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

func TestWithNormalize(t *testing.T) {
	callOptions := CallOptions{
		ZoomLevel:    16,
		ResponseType: Driver,
		Language:     Farsi,
		Normalize:    true,
	}

	WithNormalize()(&callOptions)

	if callOptions.Normalize {
		t.Fatalf("Normalize should be true")
	}
}
