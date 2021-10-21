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
func (m StatusCode) GetWorkStatus() string {
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
