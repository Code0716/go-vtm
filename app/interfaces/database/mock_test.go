package database_test

import (
	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/interfaces/database"
)

type mockAdminRepo struct {
	database.SQLHandlerInterface
	FakeCreateAdmin       func(newAdmin interface{}) error
	FakeFirst             func(value interface{}, where ...interface{}) error
	FakeGetAdminByEmail   func(*domain.AdminUser, string) error
	FakeIsAdminExist      func(tableName string, query interface{}, args ...interface{}) (bool, error)
	FakeCreateMember      func(newMember domain.Member) error
	FakeAdminMemberGetAll func(params domain.Pager)
}

func (m mockAdminRepo) Create(adminU interface{}) error {
	err := m.FakeCreateAdmin(&adminU)
	return err
}
func (m mockAdminRepo) First(value interface{}, where ...interface{}) error {
	err := m.FakeFirst(value, where...)
	return err
}
func (m mockAdminRepo) GetAdminBFyEmail(adminU *domain.AdminUser, mail string) error {
	err := m.FakeGetAdminByEmail(adminU, mail)
	return err
}

func (m mockAdminRepo) IsExist(tableName string, query interface{}, args ...interface{}) (bool, error) {
	isExist, err := m.FakeIsAdminExist(tableName, query, args...)
	return isExist, err
}

type mockMembersRepo struct {
	database.SQLHandlerInterface
	FakeAdminMemberGetAll func(params domain.Pager) ([]*domain.Member, int64, error)
	FakeCreateMember      func(m interface{}) error
	FakeIsMemberExist     func(tableName string, query interface{}, args ...interface{}) (bool, error)
}

func (m mockMembersRepo) AdminMemberGetAll(params domain.Pager) ([]*domain.Member, int64, error) {
	members, count, err := m.FakeAdminMemberGetAll(params)
	return members, count, err
}

func (m mockMembersRepo) Create(member interface{}) error {
	err := m.FakeCreateMember(&member)
	return err
}

func (m mockMembersRepo) IsExist(tableName string, query interface{}, args ...interface{}) (bool, error) {
	return m.FakeIsMemberExist(tableName, query, args...)
}
