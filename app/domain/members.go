package domain

import (
	"github.com/Code0716/go-vtm/app/gen/api"
)

// Member defines model for Member.
type Member api.Member

// TableName GORMにテーブル名を教える
func (i Member) TableName() string {
	return "members"
}

// AdminRegistMemberJSONRequestBody defines body for RegistMember for application/json ContentType.
type AdminRegistMemberJSONRequestBody api.AdminRegistMemberJSONRequestBody

// MembersResponse defines model for Members.
type MembersResponse struct {
	Members []*Member `json:"members"`
	Total   int64     `json:"total"`
}

// StatusCode member status model
type StatusCode int

const (
	// StatusCodeInit 初回登録時のstatus
	StatusCodeInit StatusCode = iota + 1
	// StatusCodeActive 稼働中のstatus
	StatusCodeActive
	// StatusCodeOther 休業中かやめた方
	StatusCodeOther
)

const (
	// StatusInit 初回登録時のstatus
	StatusInit = "init"
	// StatusActive 稼働中のstatus
	StatusActive = "active"
	// StatusOther 休業中かやめた方
	StatusOther = "other"
)

// GetMembeStatus get member status
func (m StatusCode) GetMembeStatus() string {
	switch m {
	case StatusCodeInit:
		return StatusInit
	case StatusCodeActive:
		return StatusActive
	case StatusCodeOther:
		return StatusOther
	default:
		return StatusOther
	}
}
