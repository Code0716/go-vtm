// Package database interface
package database

// import (
// 	"context"
// 	"log"

// 	"github.com/Code0716/go-vtm/app/domain"
// )

// // AttendanceRepository is user database.
// type AttendanceRepository struct {
// 	SQLHandler SQLHandlerInterface
// }

// // NewAttendance initializes attendance database.
// func NewAttendance(sqlHandler SQLHandlerInterface) *AttendanceRepository {
// 	return &AttendanceRepository{
// 		sqlHandler,
// 	}
// }

// // Timestamp to attendance table
// func (r *AttendanceRepository) Timestamp(_ context.Context, attendance domain.Attendance) error {
// 	err := r.SQLHandler.Create(&attendance).Conn.Error
// 	if err != nil {
// 		log.Print(err)
// 		return err
// 	}

// 	return nil
// }
