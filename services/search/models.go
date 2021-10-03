package search

type City struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Cetroid struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
	Description string `json:"description"`
}

type Result struct {
	PlaceID              string `json:"place_id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	StructuredFormatting struct {
		MainText      string `json:"main_text"`
		SecondaryText string `json:"secondary_text"`
	} `json:"structured_formatting"`
	Type     string `json:"type"`
	Location struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	} `json:"location"`
	Distance   float64  `json:"distance"`
	AreaLength float64  `json:"area_length"`
	AllTags    []string `json:"all_tags"`
}

type Detail struct {
	Name     string `json:"name"`
	Geometry struct {
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
	} `json:"geometry"`
}
