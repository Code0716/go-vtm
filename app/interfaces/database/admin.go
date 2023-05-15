// Package database is database Implementation
package database

import (
	"context"
	"time"

	"github.com/Code0716/go-vtm/app/domain"
)

// AdminRepository is admin database.
type AdminRepository struct {
	SQLHandler SQLHandlerInterface
}

// NewAdmin initializes admins database.
func NewAdmin(sqlHandler SQLHandlerInterface) *AdminRepository {
	return &AdminRepository{
		SQLHandler: sqlHandler,
	}
}

// GetAdminByEmail get admin user by mail address
func (r *AdminRepository) GetAdminByEmail(_ context.Context, mail string) (*domain.AdminUser, error) {
	var adminUser domain.AdminUser
	err := r.SQLHandler.First(&adminUser, domain.AdminUser{MailAddress: mail}).Conn.Error

	if err != nil {
		return nil, err
	}
	return &adminUser, nil
}

// GetAdminByUUID get admin user by  uuid
func (r *AdminRepository) GetAdminByUUID(_ context.Context, uuid string) (*domain.AdminUser, error) {
	var adminUser domain.AdminUser
	err := r.SQLHandler.First(&adminUser, domain.AdminUser{AdminId: uuid}).Conn.Error
	if err != nil {
		return nil, err
	}

	return &adminUser, nil
}

// RegistAdmin retuern error
func (r *AdminRepository) RegistAdmin(_ context.Context, adminUser domain.AdminUser) error {
	err := r.SQLHandler.Create(&adminUser)
	if err != nil {
		return err
	}

	return nil
}

// GetAllAdminUser get admin users
func (r *AdminRepository) GetAllAdminUser(_ context.Context, params domain.Pager) ([]*domain.AdminUser, int64, error) {
	adminUsers, count, err := r.SQLHandler.GetAllAdminUsers(params)
	if err != nil {
		return nil, 0, err
	}

	return adminUsers, count, err
}

// PutAdminUser put admin user
func (r *AdminRepository) PutAdminUser(_ context.Context, params domain.AdminUser) (*domain.AdminUser, error) {
	err := r.SQLHandler.Save(&params).Conn.Error
	if err != nil {
		return nil, err
	}

	return &params, err
}

// IsAdminExist check admin is already registered
func (r *AdminRepository) IsAdminExist(_ context.Context, mail string) (bool, error) {
	var adminU domain.AdminUser
	isExist, err := r.SQLHandler.IsExist(adminU.TableName(), "mail_address = ?", mail)
	if err != nil {
		return false, err
	}

	return isExist, nil
}

// DeleteAdminUser delete admin user
func (r *AdminRepository) DeleteAdminUser(_ context.Context, uuid string) (*domain.AdminUser, error) {
	var adminUser domain.AdminUser
	err := r.SQLHandler.First(&adminUser, domain.AdminUser{AdminId: uuid}).Conn.Error
	if err != nil {
		return nil, err
	}

	currentTime := time.Now()
	adminUser.DeletedAt = &currentTime
	adminUser.Status = domain.StatusCodeOther.GetWorkStatus()
	err = r.SQLHandler.Save(&adminUser).Conn.Error
	if err != nil {
		return nil, err
	}

	return &adminUser, nil
}
