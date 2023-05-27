// Package domain as Attendance
package domain

import "time"

// Attendance model
type Attendance struct {
	ID        string     `json:"id"`
	UserID    string     `json:"userId"`
	StartTime string     `json:"startTime"`
	EndTime   *string    `json:"endTime,omitempty"`
	BreakTime *time.Time `json:"breakTime,omitempty"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

// TableName GORMにテーブル名を教える
func (i Attendance) TableName() string {
	return "attendances"
}
