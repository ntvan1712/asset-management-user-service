package router

import (
	"user_service/common/middleware"
	"user_service/module/user/controller"

	"github.com/gofiber/fiber/v2"
)

const (
	UserIDParamName = "user_id"
)

func Setup(app *fiber.App) {

	userController := controller.NewUserController()

	managerRoute := app.Group("/api/managers")

	managerRoute.Use(middleware.GetAuthMiddleware().AdminAuthorityMiddleware)
	managerRoute.Get("/", userController.GetAllManagersHandler)
	managerRoute.Post("/", userController.AddManagerHandler)
	managerRoute.Patch("/:manager_id/permissions", userController.UpdateManagerPermissionsHandler)
	managerRoute.Delete("/:manager_id", userController.DeleteManagerHandler)

	employeeRoute := app.Group("/api/employees")
	employeeRoute.Use(middleware.GetAuthMiddleware().AdminAuthorityMiddleware)
	employeeRoute.Get("/", userController.SearchEmployeesByNameOrCodeHandler)
}
