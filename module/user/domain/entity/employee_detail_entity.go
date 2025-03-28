package entity

import (
	"time"
	"user_service/app_config"
	"user_service/common/proto_share"
)

type EmployeeDetailEntity struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	Code           string     `json:"code"`
	Email          string     `json:"email"`
	PhoneNumber    string     `json:"phone_number"`
	AvatarUrl      string     `json:"avatar_url"`
	Birthday       *time.Time `json:"birthday"`
	DepartmentName *string    `json:"department_name,omitempty"`
	PositionName   *string    `json:"position_name,omitempty"`
}

func EmployeeDetailEntityFromGRpc(employeeGRpc *proto_share.EmployeeRPC) EmployeeDetailEntity {
	employeeEntity := EmployeeDetailEntity{
		ID:             int(employeeGRpc.Id),
		Name:           employeeGRpc.Name,
		Code:           employeeGRpc.Code,
		Email:          employeeGRpc.Email,
		PhoneNumber:    employeeGRpc.PhoneNumber,
		AvatarUrl:      app_config.GetAppConfig().MinioConfig.GetEmployeeDataUrl(employeeGRpc.AvatarPath),
		DepartmentName: &employeeGRpc.DepartmentName,
		PositionName:   &employeeGRpc.PositionName,
	}
	if employeeGRpc.Birthday != nil {
		birthDay := employeeGRpc.Birthday.AsTime()
		employeeEntity.Birthday = &birthDay
	}
	return employeeEntity
}

func EmployeeDetailEntitiesFromGRpcs(employeeGRpcs []*proto_share.EmployeeRPC) []EmployeeDetailEntity {
	var employeesEntities []EmployeeDetailEntity

	for _, employeeGrpc := range employeeGRpcs {
		employeesEntities = append(employeesEntities, EmployeeDetailEntityFromGRpc(employeeGrpc))
	}

	return employeesEntities
}