package main

import (
	"log"
	app_config "user_service/app_config"
	cmd "user_service/cmd"
	"user_service/common/logger"
	infras "user_service/infras"
	authRouter "user_service/module/auth/router"
	userRouter "user_service/module/user/router"
)

func main() {
	cmd.Execute()
	// Read config file
	app_config.LoadAppConfig(cmd.ConfigFileName)

	if err := logger.InitAppLogger(); err != nil {
		log.Fatal("[Main] Failed to create ZapLogger", err)
	}
	defer logger.Sync()
	logger.Info("[Main] Complete ZapLogger configuration")

	employeeServiceConn := infras.GetEmployeeServiceConn()
	defer employeeServiceConn.Close()

	fiberApp := infras.GetFiberApp()

	authRouter.Setup(fiberApp)
	userRouter.Setup(fiberApp)

	fiberAppErr := fiberApp.Listen(app_config.GetAppConfig().FiberServerConfig.HttpPort)
	if fiberAppErr != nil {
		logger.Fatal("[Main] Failed to start server: " + fiberAppErr.Error())
	}

}
