package database

import "github.com/Code0716/go-vtm/app/domain"

// SQLHandlerInterface  SQLHandler
type SQLHandlerInterface interface {
	Create(value any) error
	Find(value any, where ...any) error
	First(value any, where ...any) error
	IsExist(tableName string, query any, args ...any) (bool, error)
	GetAllAdminUsers(params domain.Pager) ([]*domain.AdminUser, int64, error)
	AdminMemberGetAll(params domain.Pager) ([]*domain.Member, int64, error)
	Save(value any) error
	Delete(value any, where ...any) error
}
