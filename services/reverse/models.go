package reverse

import "reflect"

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
	Type      ResponseType `json:"type"`
	Display   string       `json:"display"`
	Language  Language     `json:"language"`
	Zoom      float64      `json:"zoom"`
	Lat       float64      `json:"lat"`
	Lon       float64      `json:"lon"`
	ID        int32        `json:"id"`
	Normalize string       `json:"normalize"`
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

type StructuralComponent struct {
	Province      string
	City          string
	County        string
	Town          string
	Village       string
	Neighbourhood string
	Suburb        string
	Locality      string
	Primary       string
	Secondary     string
	Residential   string
	ClosedWay     string
	POI           string
}

func (s StructuralComponent) NewIterator() *StructuralComponentItr {
	var components []string
	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i).String()
		if fieldValue != "" { // Only include non-empty fields
			components = append(components, fieldValue)
		}
	}
	return &StructuralComponentItr{Components: components, index: 0}
}

type StructuralComponentItr struct {
	index      int
	Components []string
}

func (s *StructuralComponentItr) Next() (string, bool) {
	if s.index >= len(s.Components) { // Corrected condition for stopping iteration
		return "", false
	}
	component := s.Components[s.index]
	s.index++
	return component, true
}

type StructuralResult struct {
	Result *StructuralComponent
	ID     int
}

const (
	province      = "province"
	city          = "city"
	county        = "county"
	town          = "town"
	village       = "village"
	neighbourhood = "neighbourhood"
	suburb        = "suburb"
	locality      = "locality"
	primary       = "primary"
	secondary     = "secondary"
	residential   = "residential"
	poi           = "poi"
)

var convertReverseTypes = map[string]struct{}{
	province:      {},
	city:          {},
	county:        {},
	town:          {},
	village:       {},
	neighbourhood: {},
	suburb:        {},
	locality:      {},
	primary:       {},
	secondary:     {},
	residential:   {},
	poi:           {},
}
