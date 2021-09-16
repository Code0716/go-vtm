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

// AdminsResponse  response body for GetAdminList
type AdminsResponse struct {
	AdminUsers []*AdminUser `json:"admin_users"`
	Total      int64        `json:"total"`
}
