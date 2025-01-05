package models

type TimeRange string

const (
	PastWeek    TimeRange = "week"
	PastMonth   TimeRange = "month"
	Past3Months TimeRange = "3months"
	PastYear    TimeRange = "year"
	NowDate     TimeRange = "date"
	CustomDate  TimeRange = "custom_date"
)

type SortOrder string

const (
	Ascending  SortOrder = "asc"
	Descending SortOrder = "desc"
)

type SortField string

const (
	SortByDateExpense SortField = "date_expense"
	SortByAmount      SortField = "amount"
	SortByCategory    SortField = "category"
	SortByDateCreate  SortField = "created_at"
)

type SortOptions struct {
	SortBy SortField
	Order  SortOrder
}

type Pagination struct {
	Cursor *int
	Limit  int
}

type ExpenseFilter struct {
	Category  *Category
	TimeRange *TimeRange
	StartDate *string
	EndDate   *string
	MinAmount *int
	MaxAmount *int
}

type ExpenseParametrs struct {
	Filter     *ExpenseFilter
	Sort       *SortOptions
	Pagination Pagination
}
