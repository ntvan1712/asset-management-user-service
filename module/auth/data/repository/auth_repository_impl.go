package repository

import (
	"context"
	"time"
	"user_service/common/enums"
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
	departmentId := int(employee.DepartmentId)
	positionId := int(employee.PositionId)
	birthDay := employee.Birthday.AsTime()
	err = a.userDataSource.InsertIfNotExists(context, model.User{
		ID:           int(employee.Id),
		Name:         employee.Name,
		Code:         employee.Code,
		DepartmentID: &departmentId,
		PositionID:   &positionId,
		Email:        employee.Email,
		PhoneNumber:  employee.PhoneNumber,
		AvatarPath:   employee.AvatarPath,
		Birthday:     &birthDay,
		CreatedAt:    time.Now().UTC(),
		RoleID:       enums.UserRoleID.Employee,
	})
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
		User:        userModel.ToEntity(),
		AccessToken: res.AccessToken,
	}, nil

}

func NewAuthRepository() AuthRepository {
	return &authRepository{
		authRemoteService: authDS.NewAuthServiceClient(infras.GetEmployeeServiceConn()),
		userDataSource:    userDS.NewUserDataSource(infras.GetDbInstance()),
	}
}
