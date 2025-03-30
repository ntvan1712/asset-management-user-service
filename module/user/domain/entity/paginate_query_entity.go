package entity

type PaginateQueryEntity struct {
	Limit int `query:"limit" validate:"required,gte=1,lte=50"`
	Page  int `query:"page" validate:"required,gte=1"`
}
