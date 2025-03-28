package model

import (
	"time"
	"user_service/app_config"
	"user_service/common/enums"
	"user_service/common/proto_share"
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

func (model *User) ToEntity() *entity.UserEntity {
	if model == nil {
		return nil
	}
	userEntity := &entity.UserEntity{
		ID:          model.ID,
		Name:        model.Name,
		Code:        model.Code,
		Email:       model.Email,
		PhoneNumber: model.PhoneNumber,
		AvatarUrl:   app_config.GetAppConfig().MinioConfig.GetEmployeeDataUrl(model.AvatarPath),
		Birthday:    model.Birthday,
		RoleID:      model.RoleID,
		Permissions: PermissionModelsToEntities(model.Permissions),
	}
	if model.Department != nil {
		userEntity.DepartmentName = &model.Department.Name
	}
	if model.Position != nil {
		userEntity.PositionName = &model.Position.Name
	}
	if model.Role != nil {
		userEntity.RoleCode = model.Role.Code
		userEntity.RoleName = model.Role.Name
	}
	return userEntity
}

func UserModelsToEntities(users []User) []entity.UserEntity {
	var userEntities []entity.UserEntity

	for _, userModel := range users {
		userEntities = append(userEntities, *userModel.ToEntity())
	}

	return userEntities
}

func NewUserFromEmployeeGRpc(employee *proto_share.EmployeeRPC) *User {
	departmentId := int(employee.DepartmentId)
	positionId := int(employee.PositionId)
	birthDay := employee.Birthday.AsTime()
	return &User{
		ID:           int(employee.Id),
		Name:         employee.Name,
		Code:         employee.Code,
		DepartmentID: &departmentId,
		PositionID:   &positionId,
		Email:        employee.Email,
		PhoneNumber:  employee.PhoneNumber,
		AvatarPath:   employee.AvatarPath,
		Birthday:     &birthDay,
		CreatedAt:    time.Now().UTC(),
		RoleID:       enums.UserRoleID.Employee,
	}
}

func (User) TableName() string {
	return infras.TableUsers
}
