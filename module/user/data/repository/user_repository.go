package repository

import (
	"context"
	"user_service/module/user/domain/entity"
)

type UserRepository interface {
	InsertIfNotExistsByID(context context.Context, userID int) (*entity.UserEntity, error)
	UpdateRoleAndPermissions(context context.Context, userID int, roleID int, permissionIDs []int) (*entity.UserEntity, error)
	FindEmployeesByNameOrCode(context context.Context, query string, page int, limit int) ([]entity.EmployeeDetailEntity, error)
	FindByID(context context.Context, userID int) (*entity.UserEntity, error)
	FindByRoleID(context context.Context, roleID int) ([]entity.UserEntity, error)
	FindAllPermissions(context context.Context) ([]entity.PermissionEntity, error)
	FindActivitiesByUserID(context context.Context, userID int, page int, limit int) ([]entity.UserActivityEntity, error)
}
