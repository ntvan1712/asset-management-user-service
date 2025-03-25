package entity

import "user_service/module/user/domain/entity"

type LoginSuccessResponseEntity struct {
	User        entity.UserEntity `json:"user"`
	AccessToken string            `json:"access_token"`
}
