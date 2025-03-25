package repository

import (
	"context"
	"user_service/infras"
	datasource "user_service/module/user/data/data_source"
	"user_service/module/user/data/model"
	"user_service/module/user/domain/entity"
)

type userRepositoryImpl struct {
	userDS datasource.UserDataSource
}

func (d *userRepositoryImpl) FindByID(context context.Context, userID int) (*entity.UserEntity, error) {
	userModel, err := d.userDS.FindByID(context, userID)
	if err != nil {
		return nil, err
	}
	userEntity := userModel.ToEntity()
	return &userEntity, nil
}

func (d *userRepositoryImpl) FindByRoleID(context context.Context, roleID int) ([]entity.UserEntity, error) {
	userModels, err := d.userDS.FindByRoleID(context, roleID)
	if err != nil {
		return nil, err
	}
	return model.UserModelsToEntities(userModels), nil
}

// func (d *userRepositoryImpl) HasRoleAndPermissionID(
// 	context context.Context,
// 	userID int,
// 	roleID int,
// 	permissionID *int,
// ) (bool, error) {
// 	return d.userDS.HasRoleAndPermissionID(context, userID, roleID, permissionID)
// }

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{
		userDS: datasource.NewUserDataSource(infras.GetDbInstance()),
	}
}
