package model

import (
	"user_service/infras"
	"user_service/module/user/domain/entity"
)

type Permission struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Code        string `gorm:"size:50" json:"code"`
	DisplayName string `gorm:"size:100" json:"display_name"`
}

func (Permission) TableName() string {
	return infras.TablePermissions
}

func (p *Permission) ToEntity() entity.PermissionEntity {
	return entity.PermissionEntity{
		Code:        p.Code,
		DisplayName: p.DisplayName,
		ID:          p.ID,
	}
}

func PermissionModelsToEntities(permissions []Permission) []entity.PermissionEntity {
	var permissionEntities []entity.PermissionEntity

	for _, perm := range permissions {
		permissionEntities = append(permissionEntities, perm.ToEntity())
	}

	return permissionEntities
}
