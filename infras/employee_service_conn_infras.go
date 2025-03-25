package infras

import (
	"user_service/app_config"
	"sync"
	"user_service/common/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var employeeServiceConn *grpc.ClientConn
var initEmployeeServiceConnOnce sync.Once

func GetEmployeeServiceConn() *grpc.ClientConn {
	initEmployeeServiceConnOnce.Do(initEmployeeServiceConn)
	return employeeServiceConn
}

func initEmployeeServiceConn() {
	config := app_config.GetAppConfig().GRPCConnectionConfig

	conn, err := grpc.NewClient(config.EmployeeService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("[EmployeeServiceConnInfras] init error", err)
	}

	employeeServiceConn = conn

	logger.Info("[EmployeeServiceConnInfras] Init completed")
}
