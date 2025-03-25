package controller

import (
	"user_service/common/error_app"
	"user_service/module/user/domain/usecase"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func (ac *UserController) GetAllManagersHandler(c *fiber.Ctx) error {
	response, err := ac.userUsecase.GetAllManagers(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(response)
}

func NewUserController() *UserController {
	return &UserController{
		userUsecase: usecase.NewUserUsecase(),
	}
}
