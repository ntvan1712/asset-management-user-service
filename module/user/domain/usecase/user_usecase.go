package usecase

import (
	"context"
	"user_service/module/user/domain/entity"
)

type UserUsecase interface {
	// InsertIfNotExists(context context.Context, userModel model.User) error
	// FindByID(context context.Context, userID int) (*entity.UserEntity, error)
	GetAllManagers(context context.Context) ([]entity.UserEntity, error)
}
