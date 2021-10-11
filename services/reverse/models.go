package reverse

// Component is the struct containing data about a single component of an address.
type Component struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
