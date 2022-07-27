package database

import (
	"context"

	"github.com/Code0716/go-vtm/app/domain"
)

// AttendanceRepository is member database.
type AttendanceRepository struct {
	SQLHandler SQLHandlerInterface
}

// NewAttendance initializes attendance database.
func NewAttendance(sqlHandler SQLHandlerInterface) *AttendanceRepository {
	return &AttendanceRepository{
		sqlHandler,
	}
}

// Timestamp to attendance table
func (r *AttendanceRepository) Timestamp(ctx context.Context, params *domain.Attendance) error {
	return nil
}
