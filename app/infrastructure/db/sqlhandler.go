package db

import (
	"gorm.io/gorm"
)

// SQLHandler is Sql Handler
type SQLHandler struct {
	Conn *gorm.DB
}

// Create db create
func (h SQLHandler) Create(value any) SQLHandler {
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
func (h SQLHandler) First(value any, where ...any) SQLHandler {
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

// // TODO: admin 下記はなくす予定。
// // interface層で抽象化したものを使う。
// // トランザクション系は後ほど。

// // IsExist check data exests
// func (h SQLHandler) IsExist(tableName string, query any, args ...any) (bool, error) {
// 	var count int64

// 	err := h.Conn.Table(tableName).
// 		Where(query, args...).
// 		Count(&count).Error
// 	if err != nil {
// 		return false, domain.WrapInternalError(err)
// 	}

// 	if count > 0 {
// 		return true, nil
// 	}
// 	return false, nil
// }

// // GetAllAdminUsers return users found by params
// func (h SQLHandler) GetAllAdminUsers(params domain.Pager) ([]*domain.AdminUser, int64, error) {
// 	query := h.Conn
// 	if params.Status != "" {
// 		query = query.Where("status = ?", params.Status)
// 	}

// 	var count int64
// 	adminUsers := make([]*domain.AdminUser, 0)
// 	err := query.
// 		Limit(params.Limit).
// 		Offset(params.Offset).
// 		Find(&adminUsers).
// 		Count(&count).
// 		Error
// 	if err != nil {
// 		return nil, 0, domain.WrapInternalError(err)
// 	}

// 	return adminUsers, count, nil
// }

// // Users

// // AdminUserGetAll return users found by params
// func (h SQLHandler) AdminUserGetAll(params domain.Pager) ([]*domain.User, int64, error) {
// 	query := h.Conn

// 	if params.Status != "" {
// 		query = query.Where("status = ?", params.Status)
// 	}

// 	users := make([]*domain.User, 0)
// 	var count int64

// 	err := query.
// 		Limit(params.Limit).
// 		Offset(params.Offset).
// 		Find(&users).
// 		Count(&count).
// 		Error
// 	if err != nil {
// 		return nil, 0, domain.WrapInternalError(err)
// 	}

// 	return users, count, nil
// }
