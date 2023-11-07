package reverse

// Component is the struct containing data about a single component of an address.
type Component struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// FrequentAddress is the struct containing data on a frequent type request.
type FrequentAddress struct {
	Address          string `json:"address"`
	EnglishAddress   string `json:"address_en"`
	Shortname        string `json:"shortname"`
	EnglishShortname string `json:"shortname_en"`
}

// Request is the struct containing data on reverse request params
type Request struct {
	Type     ResponseType `json:"type"`
	Display  string       `json:"display"`
	Language Language     `json:"language"`
	Zoom     float64      `json:"zoom"`
	Lat      float64      `json:"lat"`
	Lon      float64      `json:"lon"`
	ID       int32        `json:"id"`
}

type BatchReverseRequest struct {
	Requests []Request `json:"requests"`
}

type Result struct {
	Result Components `json:"result"`
	ID     int        `json:"id"`
}

type ResultWithDisplayName struct {
	DisplayName DisplayName `json:"result"`
	ID          int         `json:"id"`
}

type DisplayName struct {
	DisplayName string `json:"displayName"`
}

type Components struct {
	Components []Component `json:"components"`
}

type Results struct {
	Results []Result `json:"results"`
}

type ResultsWithDisplayName struct {
	Results []ResultWithDisplayName `json:"results"`
}
