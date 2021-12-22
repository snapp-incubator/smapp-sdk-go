package matrix

// Point is the type for representing a point in a map.
type Point struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// Input is the type for providing data to matrix service.
type Input struct {
	Sources []Point `json:"sources"`
	Targets []Point `json:"targets"`
}

// Output is the type of result of a request in matrix service.
type Output struct {
	SourcesToTargets [][]struct {
		// Distance is the distance of to Point s from sources and Targets in of Input
		Distance  int    `json:"distance"`
		// Time is time prediction of eta request
		Time      int    `json:"time"`
		// FromIndex is the index of one item in sources field of Input used as origin of eta
		FromIndex int    `json:"from_index"`
		// ToIndex is the index of one item in targets field of Input used as destination of eta
		ToIndex   int    `json:"to_index"`
		// Status defines if the eta for one item of matrix is successful or not.
		Status    string `json:"status"`
	} `json:"sources_to_targets"`
}
