package entity

type SearchEmployeeQuery struct {
	Query string `query:"query" validate:"required,min=3"`
	Limit int    `query:"limit" validate:"required,gte=1,lte=50"`
	Page  int    `query:"page" validate:"required,gte=1"`
}
