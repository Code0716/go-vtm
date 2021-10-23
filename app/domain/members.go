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

// UpdateMemberUserJSONBody defines parameters for UpdateMemberUser.
type UpdateMemberUserJSONBody api.UpdateMemberUserJSONBody

// MemberResponse defines model for Members.
type MemberResponse api.MemberResponse

// MembersResponse defines model for Members.
type MembersResponse struct {
	Members []*Member `json:"members"`
	Total   int64     `json:"total"`
}
