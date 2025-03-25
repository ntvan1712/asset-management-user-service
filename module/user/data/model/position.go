package model

type Position struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Code        string `gorm:"type:varchar(50)" json:"code"`
	DisplayName string `gorm:"type:varchar(100)" json:"display_name"`
}

