package repository

import (
	"context"
	"user_service/common/error_app"
	"user_service/infras"
	datasource "user_service/module/user/data/data_source"
	"user_service/module/user/data/model"
	"user_service/module/user/domain/entity"
)

type userRepositoryImpl struct {
	userDS          datasource.UserDataSource
	employeeService datasource.EmployeeServiceClient
}

// FindActivitiesByUserID implements UserRepository.
func (d *userRepositoryImpl) FindActivitiesByUserID(context context.Context, userID int, page int, limit int) ([]entity.UserActivityEntity, error) {
	activityModels, err := d.userDS.FindActivitiesByUserID(context, userID, page, limit)
	if err != nil {
		return nil, err
	}
	return model.UserActivityModelsToEntities(activityModels), nil
}

// FindAllPermissions implements UserRepository.
func (d *userRepositoryImpl) FindAllPermissions(context context.Context) ([]entity.PermissionEntity, error) {
	permModels, err := d.userDS.FindAllPermissions(context)
	if err != nil {
		return nil, err
	}
	return model.PermissionModelsToEntities(permModels), nil
}

// FindByNameOrCode implements UserRepository.
func (d *userRepositoryImpl) FindEmployeesByNameOrCode(
	context context.Context,
	query string,
	page int,
	limit int,
) ([]entity.EmployeeDetailEntity, error) {
	request := &datasource.FindByNameOrCodeRequest{
		Query: query,
		Page:  int32(page),
		Limit: int32(limit),
	}
	response, err := d.employeeService.FindByNameOrCode(context, request)
	if err != nil {
		return nil, error_app.ErrCodeFromGRpcError(err)
	}
	return entity.EmployeeDetailEntitiesFromGRpcs(response.Employees), nil
}

// InsertIfNotExistsByID implements UserRepository.
func (d *userRepositoryImpl) InsertIfNotExistsByID(context context.Context, id int) (*entity.UserEntity, error) {
	user, err := d.userDS.FindByID(context, id)

	if err == nil {
		return user.ToEntity(), nil
	}
	if err == error_app.ErrDocumentNotFound {
		userRGpc, err := d.employeeService.FindByID(context, &datasource.FindByIDRequest{Id: int32(id)})
		if err != nil {
			return nil, error_app.ErrCodeFromGRpcError(err)
		} else {
			if err := d.userDS.InsertIfNotExists(context, model.NewUserFromEmployeeGRpc(userRGpc)); err != nil {
				return nil, err
			}
			return model.NewUserFromEmployeeGRpc(userRGpc).ToEntity(), nil
		}

	}
	return nil, err

}

// UpdateRoleAndPermissions implements UserRepository.
func (d *userRepositoryImpl) UpdateRoleAndPermissions(context context.Context, userID int, roleID int, permissionIDs []int) (*entity.UserEntity, error) {
	userModel, err := d.userDS.UpdateRoleAndPermissions(context, userID, roleID, permissionIDs)
	if err != nil {
		return nil, err
	}
	return userModel.ToEntity(), nil
}

func (d *userRepositoryImpl) FindByID(context context.Context, userID int) (*entity.UserEntity, error) {
	userModel, err := d.userDS.FindByID(context, userID)
	if err != nil {
		return nil, err
	}
	return userModel.ToEntity(), nil
}

func (d *userRepositoryImpl) FindByRoleID(context context.Context, roleID int) ([]entity.UserEntity, error) {
	userModels, err := d.userDS.FindByRoleID(context, roleID)
	if err != nil {
		return nil, err
	}
	return model.UserModelsToEntities(userModels), nil
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{
		userDS:          datasource.NewUserDataSource(infras.GetDbInstance()),
		employeeService: datasource.NewEmployeeServiceClient(infras.GetEmployeeServiceConn()),
	}
}
