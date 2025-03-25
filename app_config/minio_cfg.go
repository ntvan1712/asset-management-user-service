package app_config

import (
	"fmt"
	"strings"
)

type MinioConfig struct {
	Endpoint        string `mapstructure:"Endpoint"`
	AccessKeyID     string `mapstructure:"AccessKeyID"`
	SecretAccessKey string `mapstructure:"SecretAccessKey"`
	AssetBucket     string `mapstructure:"AssetBucket"`
	EmployeeBucket  string `mapstructure:"EmployeeBucket"`
}

func (m *MinioConfig) GetEmployeeDataUrl(path string) string {
	if strings.HasPrefix(path, "/") {
		return fmt.Sprintf("%s%s", m.Endpoint, path)
	}
	return fmt.Sprintf("%s/%s", m.Endpoint, path)
}
