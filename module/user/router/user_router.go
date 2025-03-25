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
}
