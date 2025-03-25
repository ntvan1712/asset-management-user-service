package repository

import (
	"context"
	"user_service/module/auth/domain/entity"
)

type AuthRepository interface {
	Login(context context.Context, username string, password string) (*entity.LoginSuccessResponseEntity, error)
}
