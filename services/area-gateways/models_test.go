package area_gateways

import "testing"

func TestPoint_Validate(t *testing.T) {
	t.Run("invalid_lat", func(t *testing.T) {
		p := Point{
			Lat: -180,
			Lon: 50,
		}
		err := p.Validate()
		if err == nil {
			t.Fatalf("err should not be nil. lat of point is invalid")
		}
	})
	t.Run("invalid_lon", func(t *testing.T) {
		p := Point{
			Lat: 50,
			Lon: 200,
		}
		err := p.Validate()
		if err == nil {
			t.Fatalf("err should not be nil. lon of point is invalid")
		}
	})
	t.Run("valid", func(t *testing.T) {
		p := Point{
			Lat: 50,
			Lon: 50,
		}
		err := p.Validate()
		if err != nil {
			t.Fatalf("err should be nil")
		}
	})
}
