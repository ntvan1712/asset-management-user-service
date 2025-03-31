package infras

import (
	"sync"
	"user_service/app_config"
	"user_service/common/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	TableUsers          = "users"
	TablePermissions    = "permissions"
	TableDepartments    = "departments"
	TablePositions      = "positions"
	TableRoles          = "roles"
	TableUserActivities = "user_activities" 
	TableAssetTypes     = "asset_types"
	TableAssetQualities = "asset_qualities"
	TableCurrencies     = "currencies"
	TableLocations      = "locations"
)

var dbInstance *gorm.DB
var initDbInstanceOnce sync.Once

func GetDbInstance() *gorm.DB {
	initDbInstanceOnce.Do(initDb)
	return dbInstance
}

func initDb() {
	postgresConfig := app_config.GetAppConfig().PostgresConfig
	db, err := gorm.Open(postgres.Open(postgresConfig.ConnectionString))
	if err != nil {
		logger.Fatal("[PostgresInfras] init error", err)
	}
	dbInstance = db
	logger.Info("[PostgresInfras] init PostgresDB")
}
