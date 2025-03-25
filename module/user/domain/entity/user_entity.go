package entity

import (
	"time"
	"user_service/common/enums"
)

type UserEntity struct {
	ID             int                `json:"id"`
	Name           string             `json:"name"`
	Code           string             `json:"code"`
	Email          string             `json:"email"`
	PhoneNumber    string             `json:"phone_number"`
	AvatarUrl      string             `json:"avatar_url"`
	Birthday       *time.Time         `json:"birthday"`
	DepartmentName *string            `json:"department_name,omitempty"`
	PositionName   *string            `json:"position_name,omitempty"`
	RoleID         int                `json:"-"`
	RoleCode       string             `json:"role_code"`
	RoleName       string             `json:"role_name"`
	Permissions    []PermissionEntity `json:"permissions,omitempty"`
}

func (u *UserEntity) HasAuthority(userAuthority string) bool {
	if userAuthority == enums.UserAuthority.Admin {
		return u.HasAdminAuthority()
	}
	if userAuthority == enums.UserAuthority.AssetManagement {
		return u.HasAssetManagementAuthority()
	}
	if userAuthority == enums.UserAuthority.BorrowManagement {
		return u.HasBorrowManagementAuthority()
	}
	if userAuthority == enums.UserAuthority.Statistical {
		return u.HasStatisticalAuthority()
	}

	if userAuthority == enums.UserAuthority.CategoryManagement {
		return u.HasCategoryManagementAuthority()
	}
	return u.HasEmployeeAuthority()
}

func (u *UserEntity) HasAdminAuthority() bool {
	return u.RoleID == enums.UserRoleID.Admin
}

func (u *UserEntity) hasManagerAuthority(permissionID int) bool {
	if u.HasAdminAuthority() {
		return true
	}
	if u.RoleID == enums.UserRoleID.Manager {
		for _, perm := range u.Permissions {
			if perm.ID == permissionID {
				return true
			}
		}
	}
	return false
}

func (u *UserEntity) HasAssetManagementAuthority() bool {
	return u.hasManagerAuthority(enums.PermissionID.AssetManagement)

}

func (u *UserEntity) HasBorrowManagementAuthority() bool {
	return u.hasManagerAuthority(enums.PermissionID.BorrowManagement)
}

func (u *UserEntity) HasStatisticalAuthority() bool {
	return u.hasManagerAuthority(enums.PermissionID.Statistical)
}

func (u *UserEntity) HasCategoryManagementAuthority() bool {
	return u.hasManagerAuthority(enums.PermissionID.CategoryManagement)
}

func (u *UserEntity) HasEmployeeAuthority() bool {
	return true
}
