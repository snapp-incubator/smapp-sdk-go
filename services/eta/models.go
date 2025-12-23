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
	Metadata map[string]string `json:"metadata,omitempty"`
}
