package entity

type PermissionEntity struct {
	ID   int    `json:"id"`
	Code string `gorm:"size:50" json:"code"`
	Name string `gorm:"size:100" json:"name"`
}
