package interactor_test

import (
	"context"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/interfaces/repository"
)

type mockAdminRepo struct {
	repository.AdminInterface
	FakeRegistAdmin  func(ctx context.Context, params domain.AdminUser) error
	FakeIsAdminExist func(ctx context.Context, name, mail string) (bool, error)
}

func (m mockAdminRepo) RegistAdmin(ctx context.Context, params domain.AdminUser) error {
	return m.FakeRegistAdmin(ctx, params)
}

func (m mockAdminRepo) IsAdminExist(ctx context.Context, name, mail string) (bool, error) {
	return m.FakeIsAdminExist(ctx, name, mail)
}

type mockMemberRepo struct {
	repository.MembersInterface
	FakeGetAll func(ctx context.Context, params domain.Pager) ([]*domain.Member, int64, error)
}

func (m mockMemberRepo) AdminMemberGetAll(ctx context.Context, params domain.Pager) ([]*domain.Member, int64, error) {
	return m.FakeGetAll(ctx, params)
}
