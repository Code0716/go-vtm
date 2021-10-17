package interactors_test

import (
	"context"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/usecase/repositories"
)

type mockAdminRepo struct {
	repositories.AdminRepository
	FakeRegistAdmin     func(ctx context.Context, params domain.AdminUser) error
	FakeIsAdminExist    func(ctx context.Context, name, mail string) (bool, error)
	FakeGetAdminByUUID  func(ctx context.Context, uuid string) (*domain.AdminUser, error)
	FakeGetAllAdminUser func(ctx context.Context, params domain.Pager) ([]*domain.AdminUser, int64, error)
	FakeDeleteAdmin     func(ctx context.Context, uuid string) (*domain.AdminUser, error)
}

func (m mockAdminRepo) RegistAdmin(ctx context.Context, params domain.AdminUser) error {
	return m.FakeRegistAdmin(ctx, params)
}

func (m mockAdminRepo) IsAdminExist(ctx context.Context, name, mail string) (bool, error) {
	return m.FakeIsAdminExist(ctx, name, mail)
}

func (m mockAdminRepo) GetAdminByUUID(ctx context.Context, uuid string) (*domain.AdminUser, error) {
	return m.FakeGetAdminByUUID(ctx, uuid)
}

func (m mockAdminRepo) GetAllAdminUser(ctx context.Context, params domain.Pager) ([]*domain.AdminUser, int64, error) {
	return m.FakeGetAllAdminUser(ctx, params)
}

func (m mockAdminRepo) DeleteAdminUser(ctx context.Context, uuid string) (*domain.AdminUser, error) {
	return m.FakeDeleteAdmin(ctx, uuid)
}

type mockMemberRepo struct {
	repositories.MembersRepository
	FakeGetAll func(ctx context.Context, params domain.Pager) ([]*domain.Member, int64, error)
}

func (m mockMemberRepo) AdminMemberGetAll(ctx context.Context, params domain.Pager) ([]*domain.Member, int64, error) {
	return m.FakeGetAll(ctx, params)
}
