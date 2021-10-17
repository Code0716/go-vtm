package db

import (
	"errors"

	"github.com/Code0716/go-vtm/app/domain"
	"gorm.io/gorm"
)

// SQLHandler is Sql Handler
type SQLHandler struct {
	Conn *gorm.DB
}

// errorを返すのでいいのか。
// *gorm.DBで返してやればrepositoryで細かく記述できるが、、、。

// Create retuern error
func (h SQLHandler) Create(value interface{}) error {
	err := h.Conn.Create(value).Error
	if err != nil {
		return err
	}
	return nil
}

// Find gorm find
// TODO:未実装
func (h SQLHandler) Find(value interface{}, where ...interface{}) error {
	err := h.Conn.Find(value, where...).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.NewError(domain.ErrorTypeContentNotFound)
	}
	if err != nil {
		return err
	}
	return nil
}

// First gorm find
func (h SQLHandler) First(value interface{}, where ...interface{}) error {
	err := h.Conn.First(value, where...).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.NewError(domain.ErrorTypeContentNotFound)
	}
	return err
}

// Save is gorm save
// TODO:未実装
func (h *SQLHandler) Save(value interface{}) error {
	err := h.Conn.Save(value).Error
	return err
}

// Update is gorm Update
// TODO:未実装
func (h *SQLHandler) Update(column string, value interface{}) error {
	err := h.Conn.Update(column, value).Error
	return err
}

// Scan is gorm Scan
// TODO 実装未
func (h *SQLHandler) Scan(value interface{}) error {
	err := h.Conn.Scan(value).Error
	return err
}

// Exec is gorm Exec
// TODO 実装未
func (h *SQLHandler) Exec(sql string, value ...interface{}) error {
	err := h.Conn.Exec(sql, value).Error
	return err
}

// Raw is gorm Raw
// TODO 実装未
func (h *SQLHandler) Raw(sql string, values ...interface{}) *gorm.DB {
	return h.Conn.Raw(sql, values...)
}

// Delete is gorm delete
func (h *SQLHandler) Delete(value interface{}, where ...interface{}) error {
	err := h.Conn.Delete(value, where...).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.NewError(domain.ErrorTypeContentNotFound)
	}
	return err
}

// Where gorm Where
func (h *SQLHandler) Where(query interface{}, args ...interface{}) error {
	err := h.Conn.Where(query, args...).Error
	return err
}

// TODO:下記も抽象化したい

// admin

// IsExist check data exests
func (h SQLHandler) IsExist(tableName string, query interface{}, args ...interface{}) (bool, error) {
	var count int64

	err := h.Conn.Table(tableName).
		Where(query, args...).
		Count(&count).Error
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}

// GetAllAdminUsers return members found by params
func (h SQLHandler) GetAllAdminUsers(params domain.Pager) ([]*domain.AdminUser, int64, error) {
	query := h.Conn
	if params.Status != "" {
		query = query.Where("status = ?", params.Status)
	}

	var count int64
	adminUsers := make([]*domain.AdminUser, 0)
	err := query.
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&adminUsers).
		Count(&count).
		Error
	if err != nil {
		return nil, 0, err
	}

	return adminUsers, count, nil
}

// Members

// AdminMemberGetAll return members found by params
func (h SQLHandler) AdminMemberGetAll(params domain.Pager) ([]*domain.Member, int64, error) {
	query := h.Conn

	if params.Status != "" {
		query = query.Where("status = ?", params.Status)
	}

	members := make([]*domain.Member, 0)
	var count int64

	err := query.
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&members).
		Count(&count).
		Error
	if err != nil {
		return nil, 0, err
	}

	return members, count, nil
}
