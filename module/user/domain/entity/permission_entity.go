package entity

type PermissionEntity struct {
	ID          int    `json:"-"`
	Code        string `gorm:"size:50" json:"code"`
	DisplayName string `gorm:"size:100" json:"display_name"`
}
