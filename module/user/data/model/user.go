package model

import (
	"time"
	"user_service/app_config"
	"user_service/infras"
	"user_service/module/user/domain/entity"
)

type User struct {
	ID           int        `gorm:"primaryKey;autoIncrement:false" json:"id"`
	Name         string     `gorm:"type:varchar(100);not null" json:"name"`
	Code         string     `gorm:"type:varchar(50)" json:"code"`
	DepartmentID *int       `gorm:"type:int" json:"department_id"`
	PositionID   *int       `gorm:"type:int" json:"position_id"`
	Email        string     `gorm:"type:varchar(100)" json:"email"`
	PhoneNumber  string     `gorm:"type:varchar(20)" json:"phone_number"`
	AvatarPath   string     `gorm:"type:text" json:"avatar_path"`
	Birthday     *time.Time `gorm:"type:date" json:"birthday"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	RoleID       int        `gorm:"type:int" json:"role_id"`

	// Relationships
	Department  *Department  `gorm:"foreignKey:DepartmentID" json:"department,omitempty"`
	Position    *Position    `gorm:"foreignKey:PositionID" json:"position,omitempty"`
	Role        *Role        `gorm:"foreignKey:RoleID" json:"role,omitempty"`
	Permissions []Permission `gorm:"many2many:user_permissions" json:"permissions,omitempty"`
}

func (model *User) ToEntity() entity.UserEntity {
	return entity.UserEntity{
		ID:          model.ID,
		Name:        model.Name,
		Code:        model.Code,
		Email:       model.Email,
		PhoneNumber: model.PhoneNumber,
		AvatarUrl:   app_config.GetAppConfig().MinioConfig.GetEmployeeDataUrl(model.AvatarPath),
		Birthday:    model.Birthday,
		// DepartmentCode: &model.Department.Code,
		DepartmentName: &model.Department.DisplayName,
		// PositionCode:   &model.Position.Code,
		PositionName: &model.Position.DisplayName,
		RoleCode:     model.Role.Code,
		RoleName:     model.Role.DisplayName,
		RoleID:       model.RoleID,
		Permissions:  PermissionModelsToEntities(model.Permissions),
	}
}

func UserModelsToEntities(users []User) []entity.UserEntity {
	var userEntities []entity.UserEntity

	for _, userModel := range users {
		userEntities = append(userEntities, userModel.ToEntity())
	}

	return userEntities
}

func (User) TableName() string {
	return infras.TableUsers
}
