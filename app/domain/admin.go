package domain

import (
	"github.com/Code0716/go-vtm/app/gen/api"
)

// AdminUser defines model for AdminUser.
type AdminUser api.AdminUser

// TableName GORMにテーブル名を教える
func (i AdminUser) TableName() string {
	return "admin_users"
}

// RegistAdminJSONRequestBody is request body
type RegistAdminJSONRequestBody api.RegistAdminJSONRequestBody

// AdminUsersResponse  response body for GetAdminList
type AdminUsersResponse struct {
	AdminUsers []*AdminUser `json:"admin_users"`
	Total      int64        `json:"total"`
}

// AdminUserResponse defines model for AdminUserResponse.
type AdminUserResponse api.AdminUserResponse

// UpdateAdminInfoJSONRequestBody defines model for UpdateAdminInfoJSONRequestBody.
type UpdateAdminInfoJSONRequestBody api.UpdateAdminInfoJSONRequestBody

// DeleteAdminUserResponse type
type DeleteAdminUserResponse api.DeleteAdminUserResponse
