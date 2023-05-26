package domain

import "time"

type UserStatus string

const (
	UserStatusInit   UserStatus = "init"
	UserStatusActive UserStatus = "active"
	UserStatusOther  UserStatus = "other"
)

type UserRole string

const (
	UserRoleAdmin      UserRole = "admin"
	UserRoleManager    UserRole = "manager"
	UserRoleAccountant UserRole = "accountant"
	UserRoleCommon     UserRole = "common"
)

type EmploymentStatus string

const (
	EmploymentStatusAnnual  EmploymentStatus = "annual"
	EmploymentStatusMonthly EmploymentStatus = "monthly"
	EmploymentStatusHourly  EmploymentStatus = "hourly"
	EmploymentStatusDay     EmploymentStatus = "day"
	EmploymentStatusOther   EmploymentStatus = "other"
)

type User struct {
	ID               string           `json:"id"`
	UserID           string           `json:"userId"`
	Name             string           `json:"name"`
	Password         string           `json:"password"`
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
