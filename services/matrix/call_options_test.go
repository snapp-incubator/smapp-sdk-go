package matrix

import "testing"

func TestMartixEngineString(t *testing.T) {
	t.Run("test enginev1", func(t *testing.T) {
		if MatrixEngineV1.String() != "v1" {
			t.Fatal("MatrixEngineV1 is not stringified correctly")
		}
	})
	t.Run("test enginev2", func(t *testing.T) {
		if MatrixEngineV2.String() != "v2" {
			t.Fatal("MatrixEngineV1 is not stringified correctly")
		}
	})
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

func TestWithTraffic(t *testing.T) {
	callOptions := CallOptions{}

	WithTraffic()(&callOptions)

	if callOptions.UseNoTraffic == false {
		t.Fatalf("UseTraffic should not be false")
	}

	if callOptions.NoTraffic == true {
		t.Fatalf("NoTraffic should not be true")
	}
}

func TestWithEngine(t *testing.T) {
	callOptions := CallOptions{}

	WithEngine(MatrixEngineV2)(&callOptions)

	if callOptions.Engine != MatrixEngineV2 {
		t.Fatalf("Engine should be MatrixEngineV2")
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
