package router

import (
	"user_service/common/middleware"
	"user_service/module/user/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	userController := controller.NewUserController()

	managerRoute := app.Group("/api/managers")

	managerRoute.Use(middleware.GetAuthMiddleware().AdminAuthorityMiddleware)
	managerRoute.Get("/", userController.GetAllManagersHandler)
	managerRoute.Post("/", userController.AddManagerHandler)
	managerRoute.Patch("/:manager_id/permissions", userController.UpdateManagerPermissionsHandler)
	managerRoute.Get("/:manager_id/activities", userController.GetManagerActivitiesHandler)
	managerRoute.Delete("/:manager_id", userController.DeleteManagerHandler)


	managerRoute.Get("/permissions", userController.GetAllPermissionsHandler)

	employeeRoute := app.Group("/api/employees")
	employeeRoute.Use(middleware.GetAuthMiddleware().BorrowManagementAuthorityMiddleware)
	employeeRoute.Get("/", userController.SearchEmployeesByNameOrCodeHandler)
	employeeRoute.Get("/:user_id", userController.GetUserByIDHandler)

}
