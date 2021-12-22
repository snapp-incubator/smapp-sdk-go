package matrix

import "testing"

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
	callOptions := NewDefaultCallOptions(WithHeaders(map[string]string{
		"foo": "bar",
	}))

	if callOptions.Headers["foo"] != "bar" {
		t.Fatalf("Language should be Farsi")
	}
}

