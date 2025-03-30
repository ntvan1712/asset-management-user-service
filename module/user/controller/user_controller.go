package controller

import (
	"user_service/common/error_app"
	"user_service/common/logger"
	"user_service/common/validator_app"
	"user_service/module/user/domain/entity"
	"user_service/module/user/domain/usecase"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func (ac *UserController) SearchEmployeesByNameOrCodeHandler(c *fiber.Ctx) error {
	searchEmployeeQuery := new(entity.SearchEmployeeQuery)
	if err := c.QueryParser(searchEmployeeQuery); err != nil {
		logger.Error("UserController", "SearchEmployeesByNameOrCodeHandler QueryParserErr", err)
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse(err.Error()))
	}

	if err := validator_app.ValidateStruct(searchEmployeeQuery); err != nil {
		logger.Error("UserController", "SearchEmployeesByNameOrCodeHandler ValidateStructErr", err)
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	response, err := ac.userUsecase.SearchEmployeesByNameOrCode(
		c.Context(),
		searchEmployeeQuery.Query,
		searchEmployeeQuery.Page,
		searchEmployeeQuery.Limit,
	)
	if err != nil {
		if err == error_app.ErrDocumentNotFound {
			return c.Status(fiber.StatusNotFound).JSON(error_app.NotFoundErrorResponse("Không có nhân sự nào trùng khớp"))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(response)
}

func (ac *UserController) GetAllPermissionsHandler(c *fiber.Ctx) error {
	response, err := ac.userUsecase.GetAllPermissions(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(response)
}

func (ac *UserController) GetAllManagersHandler(c *fiber.Ctx) error {
	response, err := ac.userUsecase.GetAllManagers(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(response)
}

func (ac *UserController) DeleteManagerHandler(c *fiber.Ctx) error {
	managerID, err := c.ParamsInt("manager_id")
	if err != nil {
		logger.Error("UserController", "UpdateManagerHandler", err)
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse("Manager Id phải là số nguyên"))
	}

	err = ac.userUsecase.DeleteManager(c.Context(), managerID)
	if err != nil {
		if err == error_app.ErrDocumentNotFound {
			return c.Status(fiber.StatusNotFound).JSON(error_app.NotFoundErrorResponse("Không tìm thấy người dùng"))
		}
		if err == error_app.ErrPermissionDenied {
			return c.Status(fiber.StatusForbidden).JSON(error_app.NotFoundErrorResponse("Không có quyền"))

		}
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.SendString("Delete success")
}

func (ac *UserController) GetManagerActivitiesHandler(c *fiber.Ctx) error {
	managerID, err := c.ParamsInt("manager_id")
	if err != nil {
		logger.Error("UserController", "GetManagerActivitiesHandler", err)
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse("Manager Id phải là số nguyên"))
	}

	paginateQuery := new(entity.PaginateQueryEntity)
	if err := c.QueryParser(paginateQuery); err != nil {
		logger.Error("UserController", "GetManagerActivitiesHandler QueryParserErr", err)
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse(err.Error()))
	}

	if err := validator_app.ValidateStruct(paginateQuery); err != nil {
		logger.Error("UserController", "GetManagerActivitiesHandler ValidateStructErr", err)
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	response, err := ac.userUsecase.GetManagerActivities(
		c.Context(),
		managerID,
		paginateQuery.Page,
		paginateQuery.Limit,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(response)
}

func (ac *UserController) UpdateManagerPermissionsHandler(c *fiber.Ctx) error {
	type updatePermissionsRequest struct {
		PermissionIDs []int `json:"permission_ids" validate:"required,min=1,dive,oneof=1 2 3 4"`
	}
	var request *updatePermissionsRequest
	if err := c.BodyParser(&request); err != nil {
		logger.Error("UserController", "UpdateManagerHandler", err)
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse(err.Error()))
	}
	if err := validator_app.ValidateStruct(request); err != nil {
		logger.Error("UserController", "UpdateManagerHandler", err)
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	managerID, err := c.ParamsInt("manager_id")
	if err != nil {
		logger.Error("UserController", "UpdateManagerHandler", err)
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse("Manager Id phải là số nguyên"))
	}

	manager, err := ac.userUsecase.UpdateManagerPermissions(c.Context(), managerID, request.PermissionIDs)
	if err != nil {
		if err == error_app.ErrDocumentNotFound {
			return c.Status(fiber.StatusNotFound).JSON(error_app.NotFoundErrorResponse("Không tìm thấy người dùng"))
		}
		if err == error_app.ErrPermissionDenied {
			return c.Status(fiber.StatusForbidden).JSON(error_app.NotFoundErrorResponse("Không có quyền"))

		}
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(manager)
}

func (ac *UserController) AddManagerHandler(c *fiber.Ctx) error {
	var request entity.CreateManagerRequestEntity
	if err := c.BodyParser(&request); err != nil {
		logger.Error("UserController", "CreateManagerHandler", err)
		return c.Status(fiber.StatusBadRequest).JSON(error_app.BadRequestErrorResponse(err.Error()))
	}
	if err := validator_app.ValidateStruct(request); err != nil {
		logger.Error("UserController", "CreateManagerHandler", err)
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	manager, err := ac.userUsecase.AddManager(c.Context(), request)
	if err != nil {
		if err == error_app.ErrDocumentNotFound {
			return c.Status(fiber.StatusNotFound).JSON(error_app.NotFoundErrorResponse("Không tìm thấy người dùng"))
		}
		if err == error_app.ErrPermissionDenied {
			return c.Status(fiber.StatusForbidden).JSON(error_app.NotFoundErrorResponse("Không có quyền"))

		}
		return c.Status(fiber.StatusInternalServerError).JSON(error_app.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(manager)
}

func NewUserController() *UserController {
	return &UserController{
		userUsecase: usecase.NewUserUsecase(),
	}
}
