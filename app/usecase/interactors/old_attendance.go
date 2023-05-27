// Package interactors as attendnce
package interactors

// import (
// 	"context"
// 	"log"
// 	"time"

// 	"github.com/Code0716/go-vtm/app/domain"
// 	"github.com/Code0716/go-vtm/app/usecase/repositories"
// )

// // AttendanceInteractor is user interactor.
// type AttendanceInteractor struct {
// 	AttendanceRepository repositories.AttendanceRepository
// }

// // NewAttendance initializes item interactor.
// func NewAttendance(
// 	attendanceRepo repositories.AttendanceRepository,
// ) *AttendanceInteractor {
// 	return &AttendanceInteractor{
// 		AttendanceRepository: attendanceRepo,
// 	}
// }

// // Timestamp set time stamp to attendance
// func (im *AttendanceInteractor) Timestamp(ctx context.Context, params domain.TimestampJSONBody) (*domain.Attendance, error) {
// 	currentTime := time.Now()
// 	attendance := &domain.Attendance{
// 		Date:     &currentTime,
// 		UserId: &params.UserId,
// 		Status:   (*string)(&params.Status),
// 	}

// 	err := im.AttendanceRepository.Timestamp(ctx, *attendance)
// 	if err != nil {
// 		log.Printf("interactor error %v", err)
// 		return nil, err
// 	}

// 	return attendance, nil
// }
