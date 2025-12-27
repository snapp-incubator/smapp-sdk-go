package eta

// ETA is the response type of eta service
type ETA struct {
	Trip struct {
		Legs []struct {
			Time   int `json:"time"`
			Length int `json:"length"`
		} `json:"legs"`
	} `json:"trip"`
}

// Point is the type for representing a point in a map
type Point struct {
	Lat      float64           `json:"lat"`
	Lon      float64           `json:"lon"`
	Metadata map[string]string `json:"m,omitempty"`
}

// ETARequest is the request data structure for ETA service
type ETARequest struct {
	Locations         []Point           `json:"locations"`
	DepartureDateTime string            `json:"departure_date_time,omitempty"`
	Metadata          map[string]string `json:"m,omitempty"`
}
