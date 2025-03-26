package datasource

import (
	"context"
	"user_service/common/error_app"
	"user_service/module/user/data/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userDataSourceImpl struct {
	dbInstance *gorm.DB
}

// func (d *userDataSourceImpl) HasRoleAndPermissionID(
// 	context context.Context,
// 	userID int,
// 	roleID int,
// 	permissionID *int,
// ) (bool, error) {
// 	var count *int64

// 	query := `
// 		SELECT COUNT(*)
// 		FROM users u
// 		LEFT JOIN user_permissions up ON u.id = up.user_id
// 		WHERE u.id = ? AND u.role_id = ?
// 	`

// 	// Thêm điều kiện kiểm tra permission nếu permissionID không phải nil
// 	if permissionID != nil {
// 		query += " AND up.permission_id = ?"
// 		if err := d.dbInstance.Raw(query, userID, roleID, *permissionID).Scan(&count).Error; err != nil {
// 			return false, err
// 		}
// 	} else {
// 		if err := d.dbInstance.Raw(query, userID, roleID).Scan(&count).Error; err != nil {
// 			return false, err
// 		}
// 	}
// 	if count == nil {
// 		return false, error_app.ErrDocumentNotFound
// 	}

// 	return *count > 0, nil
// }

func (d *userDataSourceImpl) FindByID(context context.Context, userID int) (*model.User, error) {
	var user *model.User
	err := d.dbInstance.Preload("Department").
		Preload("Position").
		Preload("Role").
		Preload("Permissions").
		First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, error_app.ErrDocumentNotFound
	}
	return user, nil
}

func (d *userDataSourceImpl) FindByRoleID(context context.Context, roleID int) ([]model.User, error) {
	var users []model.User
	err := d.dbInstance.Preload("Department").
		Preload("Position").
		Preload("Role").
		Preload("Permissions").
		Where("role_id = ?", roleID).
		Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (d *userDataSourceImpl) InsertIfNotExists(context context.Context, userModel model.User) error {
	result := d.dbInstance.WithContext(context).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoNothing: true,                         
	}).Create(&userModel)

	return result.Error
}

func NewUserDataSource(dbInstance *gorm.DB) UserDataSource {
	return &userDataSourceImpl{
		dbInstance: dbInstance,
	}
}
