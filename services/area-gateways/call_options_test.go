package area_gateways

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
	callOptions := NewDefaultCallOptions(WithFarsiLanguage())

	if callOptions.Language != Farsi {
		t.Fatalf("Language should be Farsi")
	}
}
