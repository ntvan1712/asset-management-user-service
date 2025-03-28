package model

import (
	"user_service/infras"
	"user_service/module/user/domain/entity"
)

type Permission struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Code string `gorm:"size:50" json:"code"`
	Name string `gorm:"size:100" json:"name"`
	Description string `gorm:"type:text" json:"description"`
}

func (Permission) TableName() string {
	return infras.TablePermissions
}

func (p *Permission) ToEntity() entity.PermissionEntity {
	return entity.PermissionEntity{
		Code: p.Code,
		Name: p.Name,
		ID:   p.ID,
	}
}

func PermissionModelsToEntities(permissions []Permission) []entity.PermissionEntity {
	var permissionEntities []entity.PermissionEntity

	for _, perm := range permissions {
		permissionEntities = append(permissionEntities, perm.ToEntity())
	}

	return permissionEntities
}
