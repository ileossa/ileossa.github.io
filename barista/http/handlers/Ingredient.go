package handlers

type Ingredient struct {
	Name string `json:"name"`
}

var ingredients = []Ingredient{
	{Name: "café"},
	{Name: "eau"},
}
