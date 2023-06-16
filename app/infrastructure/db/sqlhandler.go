package db

import (
	"context"

	"gorm.io/gorm"
)

// SQLHandler is Sql Handler
type SQLHandler struct {
	Conn *gorm.DB
}

// Create db create
func (h SQLHandler) Create(_ context.Context, value any) SQLHandler {
	return SQLHandler{h.Conn.Create(value)}
}

// Update is db Update
func (h *SQLHandler) Update(column string, value any) SQLHandler {
	return SQLHandler{h.Conn.Update(column, value)}
}

// Delete is db delete
func (h *SQLHandler) Delete(value any, where ...any) SQLHandler {
	return SQLHandler{h.Conn.Delete(value, where...)}
}

// Find db find
func (h SQLHandler) Find(value any, where ...any) SQLHandler {
	return SQLHandler{h.Conn.Find(value, where...)}

}

// First db find
func (h SQLHandler) First(_ context.Context, value any, where ...any) SQLHandler {
	return SQLHandler{h.Conn.First(value, where...)}
}

// Scan is db Scan
func (h *SQLHandler) Scan(value any) SQLHandler {
	return SQLHandler{h.Conn.Scan(value)}

}

// Where db Where
func (h *SQLHandler) Where(query any, args ...any) SQLHandler {
	return SQLHandler{h.Conn.Where(query, args...)}
}

// Joins DB Joins
func (h *SQLHandler) Joins(query string, args ...any) SQLHandler {
	return SQLHandler{h.Conn.Joins(query, args...)}
}

// Group DB Group
func (h *SQLHandler) Group(name string) SQLHandler {
	return SQLHandler{h.Conn.Group(name)}
}

// Having DB Having
func (h *SQLHandler) Having(query any, args ...any) SQLHandler {
	return SQLHandler{h.Conn.Having(query, args...)}
}

// Preload DB Preload
func (h *SQLHandler) Preload(query string, args ...any) SQLHandler {
	return SQLHandler{h.Conn.Preload(query, args...)}
}

// Table DB Table
func (h *SQLHandler) Table(name string, args ...any) SQLHandler {
	return SQLHandler{h.Conn.Table(name, args...)}
}

// Save is db save
func (h *SQLHandler) Save(value any) SQLHandler {
	return SQLHandler{h.Conn.Save(value)}
}

// Pluck is db Pluck
func (h *SQLHandler) Pluck(column string, dest any) SQLHandler {
	return SQLHandler{h.Conn.Pluck(column, dest)}
}

// Error is db Error
func (h *SQLHandler) Error() error {
	return h.Conn.Error
}

// Exec is db Exec
// TODO 実装未
func (h *SQLHandler) Exec(sql string, value ...any) SQLHandler {
	return SQLHandler{h.Conn.Exec(sql, value)}
}

// Raw is db Raw
func (h *SQLHandler) Raw(sql string, values ...any) SQLHandler {
	return SQLHandler{h.Conn.Raw(sql, values...)}
}

// interface層で抽象化したものを使う。
// トランザクション系は後ほど。
