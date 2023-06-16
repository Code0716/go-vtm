package database_test

import (
	"context"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/infrastructure/db"
	"github.com/Code0716/go-vtm/app/interfaces/database"
)

type mockUsersRepo struct {
	database.SQLHandlerInterface
	FakeCreate func(c context.Context, user any) db.SQLHandler
	FakeFirst  func(c context.Context, user any) db.SQLHandler
}

func (m mockUsersRepo) CreateUser(c context.Context, u domain.User) (*domain.User, error) {

	err := m.Create(c, u).Conn.Error
	if err != nil {
		return nil, err
	}
	var user domain.User
	err = m.FakeFirst(c, user).Conn.Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m mockUsersRepo) Create(ctx context.Context, value any) db.SQLHandler {
	return m.FakeCreate(ctx, value)
}

func (m mockUsersRepo) First(ctx context.Context, value any, where ...any) db.SQLHandler {
	return m.FakeFirst(ctx, value)
}

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
