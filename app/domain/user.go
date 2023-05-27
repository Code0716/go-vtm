package domain

import "time"

// UserStatus enum
type UserStatus string

const (
	// UserStatusInit constant
	UserStatusInit UserStatus = "init"
	// UserStatusActive constant
	UserStatusActive UserStatus = "active"
	// UserStatusOther constant
	UserStatusOther UserStatus = "other"
)

// UserRole type
type UserRole string

const (
	// UserRoleAdmin constant
	UserRoleAdmin UserRole = "admin"
	// UserRoleManager constant
	UserRoleManager UserRole = "manager"
	// UserRoleAccountant constant
	UserRoleAccountant UserRole = "accountant"
	// UserRoleCommon constant
	UserRoleCommon UserRole = "common"
)

// EmploymentStatus type
type EmploymentStatus string

const (
	// EmploymentStatusAnnual constant
	EmploymentStatusAnnual EmploymentStatus = "annual"
	// EmploymentStatusMonthly constant
	EmploymentStatusMonthly EmploymentStatus = "monthly"
	// EmploymentStatusHourly constant
	EmploymentStatusHourly EmploymentStatus = "hourly"
	// EmploymentStatusDay constant
	EmploymentStatusDay EmploymentStatus = "day"
	// EmploymentStatusOther constant
	EmploymentStatusOther EmploymentStatus = "other"
)

// User is user model
type User struct {
	ID               string           `json:"id"`
	UserID           string           `json:"userId"`
	Name             string           `json:"name"`
	MailAddress      *string          `json:"mailAddress,omitempty"`
	PhoneNumber      *string          `json:"phoneNumber,omitempty"`
	Status           UserStatus       `json:"status"`
	Role             UserRole         `json:"role,omitempty"`
	EmploymentStatus EmploymentStatus `json:"employmentStatus"`
	UnitPrice        int              `json:"unitPrice"`
	DepartmentID     string           `json:"departmentId"`
	CreatedAt        time.Time        `json:"createdAt"`
	UpdatedAt        time.Time        `json:"updatedAt"`
	DeletedAt        *time.Time       `json:"deletedAt,omitempty"`
}

// TableName GORMにテーブル名を教える
func (i User) TableName() string {
	return "users"
}
