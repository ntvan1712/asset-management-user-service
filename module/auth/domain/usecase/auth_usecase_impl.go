package usecase

import (
	"context"
	"user_service/common/enums"
	authRepo "user_service/module/auth/data/repository"
	"user_service/module/auth/domain/entity"
	userRepo "user_service/module/user/data/repository"
)

type authUsecaseImpl struct {
	authRepo authRepo.AuthRepository
	userRepo userRepo.UserRepository
}

// Login implements AuthUsecase.
func (a *authUsecaseImpl) Login(context context.Context, loginRequest entity.LoginRequest) (*entity.LoginSuccessResponseEntity, error) {
	return a.authRepo.Login(context, loginRequest.UserName, loginRequest.Password)
}

// HasUserAuthority implements UserUsecase.
func (d *authUsecaseImpl) HasUserAuthority(context context.Context, userID int, userAuthority string) (bool, error) {

	// Nếu quyền cần kiểm tra là employee thì auto pass.
	if userAuthority == enums.UserAuthority.Employee {
		return true, nil
	}

	user, err := d.userRepo.FindByID(context, userID)
	if err != nil {
		return false, err
	}

	return user.HasAuthority(userAuthority), nil
}

func NewAuthUsecase() AuthUsecase {
	return &authUsecaseImpl{
		authRepo: authRepo.NewAuthRepository(),
		userRepo: userRepo.NewUserRepository(),
	}
}
