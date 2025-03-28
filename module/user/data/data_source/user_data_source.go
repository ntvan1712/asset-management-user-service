package datasource

import (
	"context"
	"user_service/module/user/data/model"
)

type UserDataSource interface {
	InsertIfNotExists(context context.Context, userModel *model.User) error
	UpdateRoleAndPermissions(context context.Context, userID int, roleID int, permissionIDs []int) (*model.User, error)
	FindByID(context context.Context, userID int) (*model.User, error)
	FindByRoleID(context context.Context, roleID int) ([]model.User, error)
	IsExistsByID(context context.Context, userID int) (bool, error)
	// HasRoleAndPermissionID(context context.Context, userID int, roleID int, permissionID *int) (bool, error)
}
