package search

// City is the struct for city candidate in city search
type City struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Centroid struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	} `json:"centroid"`
	Description string `json:"description"`
	Metadata    struct {
		CityDetail struct {
			CityId     int64 `json:"city_id"`
			DistrictId int64 `json:"district_id"`
			HexagonId  int64 `json:"hexagon_id"`
		} `json:"city_detail"`
	} `json:"metadata"`
}

// Result is the struct for search candidate of a place
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

// Detail is the struct type that specifies details of a place id in search
type Detail struct {
	Name     string `json:"name"`
	Geometry struct {
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
	} `json:"geometry"`
}
