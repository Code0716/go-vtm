package handlers_test

// import (
// 	"context"

// 	"github.com/Code0716/go-vtm/app/domain"
// 	"github.com/Code0716/go-vtm/app/registry"
// 	"github.com/Code0716/go-vtm/app/usecase/interactors"
// 	"github.com/Code0716/go-vtm/app/usecase/repositories"
// )

// type registryMock struct {
// 	registry.InteractorGetter
// 	registry.RepositoryGetter
// 	mockUserRepo
// 	mockAdminRepo
// }

// func (rm registryMock) AdminInteractor() *interactors.AdminInteractor {
// 	return interactors.NewAdmin(rm.AdminRepository())
// }

// func (rm registryMock) AdminRepository() repositories.AdminRepository {
// 	return rm.mockAdminRepo
// }

// func (rm registryMock) UsersInteractor() *interactors.UsersInteractor {
// 	return interactors.NewUsers(rm.UsersRepository())
// }

// func (rm registryMock) UsersRepository() repositories.UsersRepository {
// 	return rm.mockUserRepo
// }

// type mockAdminRepo struct {
// 	repositories.AdminRepository
// 	FakeRegistAdmin     func(ctx context.Context, params domain.AdminUser) error
// 	FakeIsAdminExist    func(ctx context.Context, mail string) (bool, error)
// 	FakeGetAdminByUUID  func(ctx context.Context, uuid string) (*domain.AdminUser, error)
// 	FakeGetAllAdminUser func(ctx context.Context, params domain.Pager) ([]*domain.AdminUser, int64, error)
// 	FakeDeleteAdminUser func(ctx context.Context, uuid string) (*domain.AdminUser, error)
// }

// func (m mockAdminRepo) RegistAdmin(ctx context.Context, params domain.AdminUser) error {
// 	return m.FakeRegistAdmin(ctx, params)
// }

// func (m mockAdminRepo) IsAdminExist(ctx context.Context, mail string) (bool, error) {
// 	return m.FakeIsAdminExist(ctx, mail)
// }

// func (m mockAdminRepo) GetAdminByUUID(ctx context.Context, uuid string) (*domain.AdminUser, error) {
// 	return m.FakeGetAdminByUUID(ctx, uuid)
// }

// func (m mockAdminRepo) GetAllAdminUser(ctx context.Context, params domain.Pager) ([]*domain.AdminUser, int64, error) {
// 	return m.FakeGetAllAdminUser(ctx, params)
// }

// func (m mockAdminRepo) DeleteAdminUser(ctx context.Context, uuid string) (*domain.AdminUser, error) {
// 	return m.FakeDeleteAdminUser(ctx, uuid)
// }

// type mockUserRepo struct {
// 	repositories.UsersRepository
// 	FakeGetAll          func(ctx context.Context, params domain.Pager) ([]*domain.User, int64, error)
// 	FakeGetUserByUUID func(ctx context.Context, uuid string) (*domain.User, error)
// }

// func (m mockUserRepo) AdminUserGetAll(ctx context.Context, params domain.Pager) ([]*domain.User, int64, error) {
// 	return m.FakeGetAll(ctx, params)
// }

// func (m mockUserRepo) GetUserByUUID(ctx context.Context, uuid string) (*domain.User, error) {
// 	return m.FakeGetUserByUUID(ctx, uuid)
// }
