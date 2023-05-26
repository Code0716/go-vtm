package database_test

// import (
// 	"github.com/Code0716/go-vtm/app/domain"
// 	"github.com/Code0716/go-vtm/app/infrastructure/db"
// 	"github.com/Code0716/go-vtm/app/interfaces/database"
// )

// type mockAdminRepo struct {
// 	database.SQLHandlerInterface
// 	FakeCreateAdmin       func(newAdmin any) db.SQLHandler
// 	FakeFirst             func(value any, where ...any) db.SQLHandler
// 	FakeGetAdminByEmail   func(*domain.AdminUser, string) error
// 	FakeIsAdminExist      func(tableName string, query any, args ...any) (bool, error)
// 	FakeCreateUser      func(newUser domain.User) error
// 	FakeAdminUserGetAll func(params domain.Pager)
// }

// func (m mockAdminRepo) Create(adminU any) error {
// 	err := m.FakeCreateAdmin(&adminU).Conn.Error
// 	return err
// }
// func (m mockAdminRepo) First(value any, where ...any) error {
// 	err := m.FakeFirst(value, where...).Conn.Error
// 	return err
// }
// func (m mockAdminRepo) GetAdminBFyEmail(adminU *domain.AdminUser, mail string) error {
// 	err := m.FakeGetAdminByEmail(adminU, mail)
// 	return err
// }

// func (m mockAdminRepo) IsExist(tableName string, query any, args ...any) (bool, error) {
// 	isExist, err := m.FakeIsAdminExist(tableName, query, args...)
// 	return isExist, err
// }

// type mockUsersRepo struct {
// 	database.SQLHandlerInterface
// 	FakeAdminUserGetAll func(params domain.Pager) ([]*domain.User, int64, error)
// 	FakeCreateUser      func(m any) error
// 	FakeIsUserExist     func(tableName string, query any, args ...any) (bool, error)
// }

// func (m mockUsersRepo) AdminUserGetAll(params domain.Pager) ([]*domain.User, int64, error) {
// 	users, count, err := m.FakeAdminUserGetAll(params)
// 	return users, count, err
// }

// func (m mockUsersRepo) Create(user any) error {
// 	err := m.FakeCreateUser(&user)
// 	return err
// }

// func (m mockUsersRepo) IsExist(tableName string, query any, args ...any) (bool, error) {
// 	return m.FakeIsUserExist(tableName, query, args...)
// }
