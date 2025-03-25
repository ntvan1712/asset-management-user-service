package usecase

import (
	"context"
	"user_service/common/enums"
	"user_service/module/user/data/repository"
	"user_service/module/user/domain/entity"
)

type userUsecaseImpl struct {
	userRepo repository.UserRepository
}

func (d *userUsecaseImpl) GetAllManagers(context context.Context) ([]entity.UserEntity, error) {
	return d.userRepo.FindByRoleID(context, enums.UserRoleID.Manager)
}

func NewUserUsecase() UserUsecase {
	return &userUsecaseImpl{
		userRepo: repository.NewUserRepository(),
	}
}
