package domain

import "github.com/Code0716/go-vtm/app/gen/api"

// Authority int
type Authority int

const (
	// AuthorityAdmin admin
	AuthorityAdmin = iota + 1
	// AuthorityManager manager
	AuthorityManager
	// AuthorityGeneral general
	AuthorityGeneral
)

// AuthorityMap is get authoryity
var AuthorityMap = map[Authority]string{
	AuthorityAdmin:   "admin",
	AuthorityManager: "manager",
	AuthorityGeneral: "general",
}

// UserStatus int
type UserStatus int

const (
	// UserStatusInit init
	UserStatusInit = iota + 1
	// UserStatusActive active
	UserStatusActive
	// UserStatusOther other
	UserStatusOther
)

// UserStatusMap is get user status
var UserStatusMap = map[UserStatus]string{
	UserStatusInit:   "init",
	UserStatusActive: "active",
	UserStatusOther:  "other",
}

// CommonSuccessResponse has a text message.
type CommonSuccessResponse api.CommonSuccessResponse

// Pager struct
type Pager struct {
	// limit params
	Limit int

	// offset param
	Offset int

	// status param
	Status string
}
