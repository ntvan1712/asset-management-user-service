package router

import (
	"user_service/common/middleware"
	"user_service/module/auth/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	authController := controller.NewAuthController()

	authRoute := app.Group("/api/auth")
	authRoute.Post("/login", authController.LoginHandler)
	authRoute.Post("/logout", middleware.GetAuthMiddleware().EmployeeAuthorityMiddleware, authController.LogoutHandler)

}
