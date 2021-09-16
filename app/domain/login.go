package domain

import "github.com/Code0716/go-vtm/app/gen/api"

// AdminLoginJSONRequestBody defines model
type AdminLoginJSONRequestBody api.AdminLoginJSONRequestBody

// MemberLoginRequest defines model
type MemberLoginRequest struct {
	Name     string `json:"name"`
	MemberID string `json:"member_id"`
}
