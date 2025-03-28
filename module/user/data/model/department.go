package model

type Department struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Code string `gorm:"type:varchar(50)" json:"code"`
	Name string `gorm:"type:varchar(100)" json:"name"`
	Description string `gorm:"type:text" json:"description"`
}
