package repository

import (
	"github.com/Code0716/go-vtm/app/domain"
)

// SQLHandlerInterface  SQLHandler
type SQLHandlerInterface interface {
	Create(value interface{}) error
	Find(value interface{}, where ...interface{}) error
	First(value interface{}, where ...interface{}) error
	IsExist(tableName string, query interface{}, args ...interface{}) (bool, error)
	GetAllAdminUsers(params domain.Pager) ([]*domain.AdminUser, int64, error)
	AdminMemberGetAll(params domain.Pager) ([]*domain.Member, int64, error)
}
