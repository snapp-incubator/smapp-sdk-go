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
