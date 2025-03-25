package controller

import (
	"user_service/common/error_app"
	"user_service/common/logger"
	"user_service/common/validator_app"
	"user_service/module/auth/domain/entity"
	"user_service/module/auth/domain/usecase"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authUsecase usecase.AuthUsecase
}

func (ac *AuthController) LoginHandler(c *fiber.Ctx) error {

	var loginRequest entity.LoginRequest
	if err := c.BodyParser(&loginRequest); err != nil {
		logger.Error("AuthController", "LoginHandler", err)
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse(err.Error()))
	}
	if err := validator_app.ValidateStruct(loginRequest); err != nil {
		logger.Error("AuthController", "LoginHandler", err)
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	response, err := ac.authUsecase.Login(c.Context(), loginRequest)
	if err != nil {
		if err == error_app.ErrBadRequest {
			return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse("Vui lòng nhập tài khoản và mật khẩu"))
		}
		if err == error_app.ErrDocumentNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(error_app.NotFoundErrorResponse("Không tìm thấy người dùng"))
		}
		if err == error_app.ErrUnauthorized {
			return c.Status(fiber.StatusUnauthorized).JSON(error_app.UnauthorizedErrorResponse("Vui lòng kiểm tra lại tài khoản và mật khẩu"))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(response)
}

func NewAuthController() *AuthController {
	return &AuthController{
		authUsecase: usecase.NewAuthUsecase(),
	}
}
