package eta

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

func TestWithNoTraffic(t *testing.T) {
	callOptions := CallOptions{}

	WithNoTraffic()(&callOptions)

	if callOptions.UseNoTraffic == false {
		t.Fatalf("UseNoTraffic should not be false")
	}

	if callOptions.NoTraffic == false {
		t.Fatalf("NoTraffic should not be false")
	}
}

func TestWithDepartureDateTime(t *testing.T) {
	callOptions := CallOptions{}

	date := "2020-10-10"

	WithDepartureDateTime(date)(&callOptions)

	if callOptions.UseDepartureDateTime == false {
		t.Fatalf("UseDepartureDateTime should not be false")
	}

	if callOptions.DepartureDateTime != date {
		t.Fatalf("DepartureDateTime should be %s but it is %s", date, callOptions.DepartureDateTime)
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
