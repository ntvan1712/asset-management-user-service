package model

import (
	"time"
	"user_service/module/user/domain/entity"
)

type UserActivity struct {
	ID            int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	ActivityType  string    `gorm:"type:varchar(50);not null" json:"activity_type"`
	DisplayName   string    `gorm:"type:varchar(255);not null" json:"display_name"`
	PriorityPoint int       `gorm:"not null" json:"priority_point"`
	UserID        int64     `gorm:"not null;index" json:"user_id"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	Description   string    `gorm:"type:text" json:"description"`
}

func (u *UserActivity) ToEntity() entity.UserActivityEntity {
	return entity.UserActivityEntity{
		ID:            u.ID,
		ActivityType:  u.ActivityType,
		DisplayName:   u.DisplayName,
		PriorityPoint: u.PriorityPoint,
		UserID:        u.UserID,
		CreatedAt:     u.CreatedAt,
		Description:   u.Description,
	}
}

func UserActivityModelsToEntities(activities []UserActivity) []entity.UserActivityEntity {
	var activityEntities []entity.UserActivityEntity

	for _, perm := range activities {
		activityEntities = append(activityEntities, perm.ToEntity())
	}

	return activityEntities
}

func (UserActivity) TableName() string {
	return "user_activities"
}
