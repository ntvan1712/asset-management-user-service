package usecase

import (
	"context"
	"user_service/module/user/domain/entity"
)

type UserUsecase interface {
	// InsertIfNotExists(context context.Context, userModel model.User) error
	// FindByID(context context.Context, userID int) (*entity.UserEntity, error)
	GetAllManagers(context context.Context) ([]entity.UserEntity, error)
	AddManager(context context.Context, createManagerRequest entity.CreateManagerRequestEntity) (*entity.UserEntity, error)
	DeleteManager(context context.Context, managerID int) error
	UpdateManagerPermissions(context context.Context, managerID int, permissionIDs []int) (*entity.UserEntity, error)

	SearchEmployeesByNameOrCode(context context.Context, query string, page int, limit int) ([]entity.EmployeeDetailEntity, error)
}
