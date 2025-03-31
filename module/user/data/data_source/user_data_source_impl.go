package datasource

import (
	"context"
	"errors"
	"user_service/common/error_app"
	"user_service/module/user/data/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userDataSourceImpl struct {
	dbInstance *gorm.DB
}

// FindActivitiesByUserID implements UserDataSource.
func (d *userDataSourceImpl) FindActivitiesByUserID(
	context context.Context,
	userID int,
	page int,
	limit int,
) ([]model.UserActivity, error) {
	var activities []model.UserActivity

	offset := (page - 1) * limit

	result := d.dbInstance.Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&activities)

	if result.Error != nil {
		return nil, result.Error
	}

	return activities, nil
}

// FindAllPermissions implements UserDataSource.
func (d *userDataSourceImpl) FindAllPermissions(context context.Context) ([]model.Permission, error) {
	var permissions []model.Permission
	err := d.dbInstance.WithContext(context).Find(&permissions).Error
	if err != nil {
		return nil, err
	}

	return permissions, nil
}

// UpdateRoleAndPermissions implements UserDataSource.
func (d *userDataSourceImpl) UpdateRoleAndPermissions(
	context context.Context,
	userID int,
	roleID int,
	permissionIDs []int,
) (*model.User, error) {
	tx := d.dbInstance.WithContext(context).Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var user *model.User
	if err := tx.First(&user, userID).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error_app.ErrDocumentNotFound
		}
		return nil, err
	}

	if err := tx.Model(user).
		Where("id = ?", userID).
		Update("role_id", roleID).Error; err != nil {

		tx.Rollback()
		return nil, err
	}

	var permissions []model.Permission
	if err := tx.Where("id IN ?", permissionIDs).Find(&permissions).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Model(&user).Association("Permissions").Replace(&permissions); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	if err := d.dbInstance.WithContext(context).
		Preload("Role").
		Preload("Department").
		Preload("Position").
		Preload("Permissions").
		First(&user, userID).Error; err != nil {
		return nil, err
	}

	return user, nil

}

func (d *userDataSourceImpl) IsExistsByID(context context.Context, userID int) (bool, error) {
	var exists bool

	err := d.dbInstance.WithContext(context).Model(&model.User{}).Select("1").Where("id = ?", userID).Scan(&exists).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, error_app.ErrDocumentNotFound
		}
		return false, err
	}

	return exists, nil
}
func (d *userDataSourceImpl) FindByID(context context.Context, userID int) (*model.User, error) {
	var user *model.User
	err := d.dbInstance.WithContext(context).Preload("Department").
		Preload("Position").
		Preload("Role").
		Preload("Permissions").
		First(&user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error_app.ErrDocumentNotFound
		}
		return nil, err
	}
	if user == nil {
		return nil, error_app.ErrDocumentNotFound
	}
	return user, nil
}

func (d *userDataSourceImpl) FindByRoleID(context context.Context, roleID int) ([]model.User, error) {
	var users []model.User
	err := d.dbInstance.WithContext(context).Preload("Department").
		Preload("Position").
		Preload("Role").
		Preload("Permissions").
		Find(&users, "role_id = ?", roleID).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error_app.ErrDocumentNotFound
		}
		return nil, err
	}

	return users, nil
}

func (d *userDataSourceImpl) InsertIfNotExists(context context.Context, userModel *model.User) error {
	result := d.dbInstance.WithContext(context).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoNothing: true,
	}).Create(&userModel)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return error_app.ErrDocumentNotFound
	}

	return result.Error
}

func NewUserDataSource(dbInstance *gorm.DB) UserDataSource {
	return &userDataSourceImpl{
		dbInstance: dbInstance,
	}
}
