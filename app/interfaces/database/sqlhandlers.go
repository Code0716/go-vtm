package database

import (
	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/infrastructure/db"
)

// SQLHandlerInterface  SQLHandler
type SQLHandlerInterface interface {
	Create(value any) db.SQLHandler
	Find(value any, where ...any) db.SQLHandler
	First(value any, where ...any) db.SQLHandler
	Save(value any) db.SQLHandler
	Delete(value any, where ...any) db.SQLHandler
	Where(query any, args ...any) db.SQLHandler
	Joins(query string, args ...any) db.SQLHandler
	Group(name string) db.SQLHandler
	Having(query any, args ...any) db.SQLHandler
	Pluck(column string, dest any) db.SQLHandler
	Error() error
	// TODO:下記は今後なくす予定
	IsExist(tableName string, query any, args ...any) (bool, error)
	GetAllAdminUsers(params domain.Pager) ([]*domain.AdminUser, int64, error)
	AdminMemberGetAll(params domain.Pager) ([]*domain.Member, int64, error)
}
