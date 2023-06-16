package database

import (
	"context"

	"github.com/Code0716/go-vtm/app/infrastructure/db"
)

// SQLHandlerInterface  SQLHandler
type SQLHandlerInterface interface {
	Create(ctx context.Context, value any) db.SQLHandler
	Update(column string, value any) db.SQLHandler
	Delete(value any, where ...any) db.SQLHandler
	Find(value any, where ...any) db.SQLHandler
	First(ctx context.Context, value any, where ...any) db.SQLHandler
	Save(value any) db.SQLHandler
	Where(query any, args ...any) db.SQLHandler
	Joins(query string, args ...any) db.SQLHandler
	Group(name string) db.SQLHandler
	Having(query any, args ...any) db.SQLHandler
	Preload(query string, args ...any) db.SQLHandler
	Pluck(column string, dest any) db.SQLHandler
	Error() error
}
