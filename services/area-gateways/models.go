package area_gateways

import "errors"

// Area is the response type of are-gateways service in golang
type Area struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
	Gates       []Gate        `json:"gates"`
}

// Gate is the type for each gate existing in Area
type Gate struct {
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

// Point is a helper type for sending request body to area-gateways service
type Point struct {
	Lat float64 `json:"lat" binding:"required"`
	Lon float64 `json:"lon" binding:"required"`
}

func (p Point) Validate() error {
	if p.Lat < -90 || p.Lat > 90 {
		return errors.New("point latitude is not valid")
	}
	if p.Lon < -180 || p.Lon > 180 {
		return errors.New("point longitude is not valid")
	}
	return nil
}
