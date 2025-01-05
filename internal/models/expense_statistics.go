package models

type ExpenseStats struct {
	TotalAmount   int                        `json:"total_amount"`
	TotalCount    int                        `json:"total_count"`
	HighestAmount int                        `json:"highest_amount"`
	LowestAmount  int                        `json:"lowest_amount"`
	AverageAmount float64                    `json:"average_amount"`
	Categories    map[Category]CategoryStats `json:"categories"`
}

type CategoryStats struct {
	TotalAmount   int     `json:"total_amount"`
	TotalCount    int     `json:"total_count"`
	HighestAmount int     `json:"highest_amount"`
	LowestAmount  int     `json:"lowest_amount"`
	AverageAmount float64 `json:"average_amount"`
}

type ExpenseStatsResponse struct {
	Data *ExpenseStats `json:"data"`
}
