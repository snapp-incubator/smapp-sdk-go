package eta

import "testing"

func TestMartixEngineString(t *testing.T) {
	t.Run("test engine v1", func(t *testing.T) {
		if EtaEngineV1.String() != "v1" {
			t.Fatal("EtaEngineV1 is not stringified correctly")
		}
	})
	t.Run("test engine v2", func(t *testing.T) {
		if EtaEngineV2.String() != "v2" {
			t.Fatal("EtaEngineV2 is not stringified correctly")
		}
	})
	t.Run("test engine nostradamus", func(t *testing.T) {
		if EtaEngineNostradamus.String() != "nostradamus" {
			t.Fatal("EtaEngineNostradamus is not stringified correctly")
		}
	})
	t.Run("test engine ocelot", func(t *testing.T) {
		if EtaEngineOcelot.String() != "ocelot" {
			t.Fatal("EtaEngineOcelot is not stringified correctly")
		}
	})
	t.Run("test engine orca", func(t *testing.T) {
		if EtaEngineOrca.String() != "orca" {
			t.Fatal("EtaEngineOrca is not stringified correctly")
		}
	})
	t.Run("test engine orca_ch", func(t *testing.T) {
		if EtaEngineOrcaCh.String() != "orca_ch" {
			t.Fatal("EtaEngineOrcaCh is not stringified correctly")
		}
	})
}

func TestWithEngine(t *testing.T) {
	callOptions := CallOptions{}

	WithEngine(EtaEngineV2)(&callOptions)

	if callOptions.Engine != EtaEngineV2 {
		t.Fatalf("Engine should be MatrixEngineV2")
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
		t.Fatalf("UseNoTraffic should not be false")
	}

	if callOptions.NoTraffic == true {
		t.Fatalf("NoTraffic should not be true")
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
