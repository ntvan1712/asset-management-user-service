package entity

type CreateManagerRequestEntity struct {
	ManagerID     int   `json:"manager_id" validate:"required,gt=0"`
	PermissionIDs []int `json:"permission_ids" validate:"required,min=1,dive,oneof=1 2 3 4"`
}
