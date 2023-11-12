package recipes

// Represents a recipe
type Recipe struct {
	Name        string        `json:"name"`
	Ingredients []Ingredients `json:"ingredients"`
}

// Represents individual ingredients
type Ingredients struct {
	Name string `json:"name"`
}
