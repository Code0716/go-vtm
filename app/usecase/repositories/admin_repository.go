// Package repositories db interfaces
package repositories

import (
	"context"

	"github.com/Code0716/go-vtm/app/domain"
)

// AdminRepository  is data access methods to admin.
type AdminRepository interface {
	RegistAdmin(ctx context.Context, params domain.AdminUser) error
	GetAllAdminUser(ctx context.Context, params domain.Pager) ([]*domain.AdminUser, int64, error)
	GetAdminByEmail(ctx context.Context, mail string) (*domain.AdminUser, error)
	GetAdminByUUID(ctx context.Context, uuid string) (*domain.AdminUser, error)
	IsAdminExist(ctx context.Context, mail string) (bool, error)
	PutAdminUser(ctx context.Context, params domain.AdminUser) (*domain.AdminUser, error)
	DeleteAdminUser(ctx context.Context, uuid string) (*domain.AdminUser, error)
}
