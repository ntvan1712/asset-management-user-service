package usecase

import (
	"context"
	"user_service/common/enums"
	"user_service/common/error_app"
	"user_service/module/user/data/repository"
	"user_service/module/user/domain/entity"
)

type userUsecaseImpl struct {
	userRepo repository.UserRepository
}

// GetManagerActivities implements UserUsecase.
func (d *userUsecaseImpl) GetManagerActivities(context context.Context, userID int, page int, limit int) ([]entity.UserActivityEntity, error) {
	return d.userRepo.FindActivitiesByUserID(context, userID, page, limit)
}

// GetAllPermissions implements UserUsecase.
func (d *userUsecaseImpl) GetAllPermissions(context context.Context) ([]entity.PermissionEntity, error) {
	return d.userRepo.FindAllPermissions(context)
}

// DeleteManager implements UserUsecase.
func (d *userUsecaseImpl) DeleteManager(context context.Context, managerId int) error {
	if _, err := d.userRepo.UpdateRoleAndPermissions(context, managerId, enums.UserRoleID.Employee, []int{}); err != nil {
		return err
	}
	return nil
}

// UpdateManager implements UserUsecase.
func (d *userUsecaseImpl) UpdateManagerPermissions(context context.Context, managerID int, permissionIDs []int) (*entity.UserEntity, error) {
	return d.userRepo.UpdateRoleAndPermissions(
		context,
		managerID,
		enums.UserRoleID.Manager,
		permissionIDs,
	)
}

// SearchEmployeesByNameOrCode implements UserUsecase.
func (d *userUsecaseImpl) SearchEmployeesByNameOrCode(context context.Context, query string, page int, limit int) ([]entity.EmployeeDetailEntity, error) {
	return d.userRepo.FindEmployeesByNameOrCode(context, query, page, limit)
}

// AddManager implements UserUsecase.
func (d *userUsecaseImpl) AddManager(context context.Context, request entity.CreateManagerRequestEntity) (*entity.UserEntity, error) {
	user, err := d.userRepo.InsertIfNotExistsByID(context, request.ManagerID)
	if err != nil {
		return nil, err
	}
	if user.RoleID != enums.UserRoleID.Employee {
		return nil, error_app.ErrPermissionDenied
	}
	return d.userRepo.UpdateRoleAndPermissions(context, request.ManagerID, enums.UserRoleID.Manager, request.PermissionIDs)
}

func (d *userUsecaseImpl) GetAllManagers(context context.Context) ([]entity.UserEntity, error) {
	return d.userRepo.FindByRoleID(context, enums.UserRoleID.Manager)
}

func NewUserUsecase() UserUsecase {
	return &userUsecaseImpl{
		userRepo: repository.NewUserRepository(),
	}
}
