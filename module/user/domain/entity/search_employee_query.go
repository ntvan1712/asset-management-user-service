package entity

type SearchEmployeeQuery struct {
	Query string `query:"query" validate:"required,min=3"`
	PaginateQueryEntity
}
