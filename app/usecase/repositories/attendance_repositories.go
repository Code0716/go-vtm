package repositories

import (
	"context"

	"github.com/Code0716/go-vtm/app/domain"
)

// AttendanceRepository  is data access methods attendance.
type AttendanceRepository interface {
	Timestamp(context.Context, domain.Attendance) error
}
