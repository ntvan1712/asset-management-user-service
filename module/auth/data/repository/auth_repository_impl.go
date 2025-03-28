package repository

import (
	"context"
	"user_service/common/error_app"
	infras "user_service/infras"
	authDS "user_service/module/auth/data/data_source"
	"user_service/module/auth/domain/entity"
	userDS "user_service/module/user/data/data_source"
	"user_service/module/user/data/model"
)

type authRepository struct {
	authRemoteService authDS.AuthServiceClient
	userDataSource    userDS.UserDataSource
}

// Login implements AuthRepository.
func (a *authRepository) Login(context context.Context, username string, password string) (*entity.LoginSuccessResponseEntity, error) {
	loginRequest := authDS.LoginRequestRPC{
		Username: username,
		Password: password,
	}

	res, err := a.authRemoteService.Login(context, &loginRequest)
	if err != nil {
		return nil, error_app.ErrCodeFromGRpcError(err)
	}

	employee := res.EmployeeDetail
	err = a.userDataSource.InsertIfNotExists(context, model.NewUserFromEmployeeGRpc(res.EmployeeDetail))
	if err != nil {
		return nil, err
	}

	userModel, err := a.userDataSource.FindByID(context, int(employee.Id))
	if err != nil {
		return nil, err
	}
	if userModel == nil {
		return nil, error_app.ErrDocumentNotFound
	}

	return &entity.LoginSuccessResponseEntity{
		User:        *userModel.ToEntity(),
		AccessToken: res.AccessToken,
	}, nil

}

func NewAuthRepository() AuthRepository {
	return &authRepository{
		authRemoteService: authDS.NewAuthServiceClient(infras.GetEmployeeServiceConn()),
		userDataSource:    userDS.NewUserDataSource(infras.GetDbInstance()),
	}
}
