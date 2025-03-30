package entity

import "time"

type UserActivityEntity struct {
	ID            int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	ActivityType  string    `gorm:"type:varchar(50);not null" json:"activity_type"`
	DisplayName   string    `gorm:"type:varchar(255);not null" json:"display_name"`
	PriorityPoint int       `gorm:"not null" json:"priority_point"`
	UserID        int64     `gorm:"not null;index" json:"user_id"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	Description   string    `gorm:"type:text" json:"description"`
}