package handlers_test

import (
	"context"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/interactor"
	"github.com/Code0716/go-vtm/app/interfaces/repository"
	"github.com/Code0716/go-vtm/app/registry"
)

type registryMock struct {
	registry.InteractorGetter
	registry.RepositoryGetter
	mockMemberRepo
	mockAdminRepo
}

func (rm registryMock) AdminInteractor() *interactor.AdminInteractor {
	return interactor.NewAdmin(rm.AdminRepository())
}

func (rm registryMock) AdminRepository() repository.AdminInterface {
	return rm.mockAdminRepo
}

func (rm registryMock) MembersInteractor() *interactor.MembersInteractor {
	return interactor.NewMembers(rm.MembersRepository())
}

func (rm registryMock) MembersRepository() repository.MembersInterface {
	return rm.mockMemberRepo
}

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
