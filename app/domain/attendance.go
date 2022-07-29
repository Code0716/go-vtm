package domain

import "github.com/Code0716/go-vtm/app/gen/api"

// Attendance defines model for Attendance
type Attendance api.Attendance

// TimestampJSONBody model
type TimestampJSONBody api.TimestampJSONBody

// TableName GORMにテーブル名を教える
func (i Attendance) TableName() string {
	return "attendance"
}
