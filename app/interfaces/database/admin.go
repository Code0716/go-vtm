package database

import (
	"context"

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
func (r *AdminRepository) GetAdminByEmail(ctx context.Context, mail string) (*domain.AdminUser, error) {
	var adminUser domain.AdminUser
	err := r.SQLHandler.First(&adminUser, domain.AdminUser{MailAddress: mail})

	if err != nil {
		return nil, err
	}
	return &adminUser, nil
}

// GetAdminByUUID get admin user by  uuid
func (r *AdminRepository) GetAdminByUUID(ctx context.Context, uuid string) (*domain.AdminUser, error) {
	var adminUser domain.AdminUser
	err := r.SQLHandler.First(&adminUser, domain.AdminUser{AdminId: uuid})
	if err != nil {
		return nil, err
	}

	return &adminUser, nil
}

// RegistAdmin retuern error
func (r *AdminRepository) RegistAdmin(ctx context.Context, adminUser domain.AdminUser) error {
	err := r.SQLHandler.Create(&adminUser)
	if err != nil {
		return err
	}
	return nil
}

// GetAllAdminUser get admin users
func (r *AdminRepository) GetAllAdminUser(ctx context.Context, params domain.Pager) ([]*domain.AdminUser, int64, error) {
	adminUsers, count, err := r.SQLHandler.GetAllAdminUsers(params)
	if err != nil {
		return nil, 0, err
	}
	return adminUsers, count, err
}

// IsAdminExist check admin is already registered
func (r *AdminRepository) IsAdminExist(ctx context.Context, name, mail string) (bool, error) {
	var adminU domain.AdminUser
	isExist, err := r.SQLHandler.IsExist(adminU.TableName(), "name = ? OR mail_address = ?", name, mail)
	if err != nil {
		return false, err
	}
	return isExist, nil
}

// DeleteAdminUser delete admin user
func (r *AdminRepository) DeleteAdminUser(ctx context.Context, uuid string) (*domain.AdminUser, error) {
	var adminUser domain.AdminUser
	err := r.SQLHandler.First(&adminUser, domain.AdminUser{AdminId: uuid})
	if err != nil {
		return nil, err
	}

	err = r.SQLHandler.Delete(&adminUser, domain.AdminUser{AdminId: uuid})
	if err != nil {
		return nil, err
	}

	return &adminUser, nil
}
