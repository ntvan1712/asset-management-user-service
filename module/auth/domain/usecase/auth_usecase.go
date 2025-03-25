package usecase

import (
	"context"
	"user_service/module/auth/domain/entity"
)

type AuthUsecase interface {
	Login(context context.Context, loginRequest entity.LoginRequest) (*entity.LoginSuccessResponseEntity, error)
	HasUserAuthority(context context.Context, userID int, userAuthority string) (bool, error)
}
