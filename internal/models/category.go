package models

type Category string

const (
	Groceries   Category = "Groceries"
	Leisure     Category = "Leisure"
	Electronics Category = "Electronics"
	Utilities   Category = "Utilities"
	Clothing    Category = "Clothing"
	Health      Category = "Health"
	Transport   Category = "Transport"
	Education   Category = "Education"
	Credits     Category = "Credits"
	Others      Category = "Others"
)

type AllCategories struct {
	Categories []Category `json:"categories"`
}

type CategoryResponse struct {
	Data *AllCategories `json:"data"`
}
