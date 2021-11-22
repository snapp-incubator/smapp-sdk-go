package locate

// Point is the type for representing a point in a map
type Point struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// LocatedPoint is the type to represent a located point by routing engines.
type LocatedPoint struct {
	Point Point `json:"point"`
}

// Result is the result type of locate service
type Result struct {
	Input         Point          `json:"input"`
	LocatedPoints []LocatedPoint `json:"snapped_points"`
}
