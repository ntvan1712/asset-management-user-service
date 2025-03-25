package repository

import (
	"context"
	"user_service/module/user/domain/entity"
)

type UserRepository interface {
	// InsertIfNotExists(context context.Context, userModel model.User) error
	FindByID(context context.Context, userID int) (*entity.UserEntity, error)
	FindByRoleID(context context.Context, roleID int) ([]entity.UserEntity, error)
	// HasRoleAndPermissionID(context context.Context, userID int, roleID int, permissionID *int) (bool, error)

}
