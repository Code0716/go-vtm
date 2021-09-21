// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package api

import (
	"time"
)

const (
	SecurityScopes = "security.Scopes"
)

// AdminUser
type AdminUser struct {
	// admin_id
	AdminId string `json:"admin_id"`

	// 権限
	Authority string `json:"authority"`

	// 登録日
	CreatedAt time.Time `json:"created_at"`

	// 削除日
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	// ID
	Id int64 `json:"id"`

	// mail address
	MailAddress string `json:"mail_address"`

	// name
	Name string `json:"name"`

	// password
	Password string `json:"password"`

	// status
	Status string `json:"status"`

	// 更新日
	UpdatedAt time.Time `json:"updated_at"`
}

// AdminUserResponse
type AdminUserResponse struct {
	// admin_id
	AdminId string `json:"admin_id"`

	// 権限
	Authority string `json:"authority"`

	// 登録日
	CreatedAt time.Time `json:"created_at"`

	// 削除日
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	// ID
	Id int64 `json:"id"`

	// mail address
	MailAddress string `json:"mail_address"`

	// name
	Name string `json:"name"`

	// status
	Status string `json:"status"`

	// 更新日
	UpdatedAt time.Time `json:"updated_at"`
}

// ログイン時のレスポンス
type AuthenticationResponse struct {
	// message
	Message string `json:"message"`

	// token
	Token *string `json:"token,omitempty"`
}

// 更新、登録、削除など成功した際の汎用レスポンス
type CommonSuccessResponse struct {
	// message
	Message string `json:"message"`
}

// エラー
type Error struct {
	// エラーコード
	Code string `json:"code"`

	// エラー内容
	Message string `json:"message"`

	// status
	Status *int64 `json:"status,omitempty"`
}

// エラーレスポンス
type ErrorResponse struct {
	// エラー一覧
	Errors []Error `json:"errors"`
}

// Member
type Member struct {
	// 登録日
	CreatedAt time.Time `json:"created_at"`

	// 削除日
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	// 時間単価
	HourlyPrice *int64 `json:"hourly_price,omitempty"`

	// ID
	Id int64 `json:"id"`

	// member_id
	MemberId string `json:"member_id"`

	// name
	Name string `json:"name"`

	// password
	Password *string `json:"password,omitempty"`

	// phone
	PhoneNumber string `json:"phone_number"`

	// status
	Status string `json:"status"`

	// 更新日
	UpdatedAt time.Time `json:"updated_at"`
}

// LimitParam defines model for LimitParam.
type LimitParam int

// OffsetParam defines model for OffsetParam.
type OffsetParam int

// StatusParam defines model for StatusParam.
type StatusParam string

// エラーレスポンス
type InternalServerError ErrorResponse

// GetAdminListParams defines parameters for GetAdminList.
type GetAdminListParams struct {
	// limit params
	Limit *LimitParam `json:"limit,omitempty"`

	// offset param
	Offset *OffsetParam `json:"offset,omitempty"`

	// status param
	Status *StatusParam `json:"status,omitempty"`
}

// AdminLoginJSONBody defines parameters for AdminLogin.
type AdminLoginJSONBody struct {
	// admin name
	MailAddress string `json:"mail_address"`

	// admin password
	Password string `json:"password"`
}

// AdminGetMemberListParams defines parameters for AdminGetMemberList.
type AdminGetMemberListParams struct {
	// limit params
	Limit *LimitParam `json:"limit,omitempty"`

	// offset param
	Offset *OffsetParam `json:"offset,omitempty"`

	// status param
	Status *StatusParam `json:"status,omitempty"`
}

// AdminRegistMemberJSONBody defines parameters for AdminRegistMember.
type AdminRegistMemberJSONBody struct {
	// member name
	Name string `json:"name"`

	// member phone number
	PhoneNumber string `json:"phone_number"`
}

// RegistAdminJSONBody defines parameters for RegistAdmin.
type RegistAdminJSONBody struct {
	// admin mail address
	MailAddress string `json:"mail_address"`

	// admin name
	Name string `json:"name"`

	// admin password
	Password string `json:"password"`
}

// AdminLoginJSONRequestBody defines body for AdminLogin for application/json ContentType.
type AdminLoginJSONRequestBody AdminLoginJSONBody

// AdminRegistMemberJSONRequestBody defines body for AdminRegistMember for application/json ContentType.
type AdminRegistMemberJSONRequestBody AdminRegistMemberJSONBody

// RegistAdminJSONRequestBody defines body for RegistAdmin for application/json ContentType.
type RegistAdminJSONRequestBody RegistAdminJSONBody

