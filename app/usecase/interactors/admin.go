package interactors

import (
	"context"
	"time"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/usecase/repositories"
	"github.com/Code0716/go-vtm/app/util"
)

// AdminInteractor is admin interactor.
type AdminInteractor struct {
	AdminRepository repositories.AdminRepository
}

// NewAdmin initializes item interactor.
func NewAdmin(
	adminsRepo repositories.AdminRepository,
) *AdminInteractor {
	return &AdminInteractor{
		AdminRepository: adminsRepo,
	}
}

// GetAdminJwtByEmail returns jwt
// ia: admin interactor
func (ia *AdminInteractor) GetAdminJwtByEmail(ctx context.Context, params domain.LoginJSONRequestBody) (*string, error) {
	env := util.Env()

	adminUser, err := ia.AdminRepository.GetAdminByEmail(ctx, params.MailAddress)
	if err != nil {
		return nil, err
	}

	if !util.CheckHush(adminUser.Password, params.Password) {
		return nil, domain.NewError(domain.ErrorTypePasswordOrEmailValidationFailed)
	}
	token, err := util.GetAdminNewToken(*adminUser, env.Signingkey)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

// GetAdminByUUID returns adminUser
// ia: admin interactor
func (ia *AdminInteractor) GetAdminByUUID(ctx context.Context, uuid string) (*domain.AdminUser, error) {
	adminUser, err := ia.AdminRepository.GetAdminByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return adminUser, nil
}

// RegistAdmin returns member list
// ia: admin interactor
func (ia *AdminInteractor) RegistAdmin(ctx context.Context, params domain.RegistAdminJSONRequestBody) error {
	hash, err := util.GetHush(params.Password)
	if err != nil {
		return err
	}
	currentTime := time.Now()

	registAdmin := domain.AdminUser{
		Name:        params.Name,
		Permission:  domain.PermissionMap[domain.PermissionAdmin],
		AdminId:     util.UUIDGenerator(),
		MailAddress: params.MailAddress,
		Status:      domain.UserStatusMap[domain.UserStatusInit],
		Password:    hash,
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}

	err = ia.AdminRepository.RegistAdmin(ctx, registAdmin)
	if err != nil {
		return err
	}

	return nil
}

// GetAdminList get admin list
func (ia *AdminInteractor) GetAdminList(ctx context.Context, params domain.Pager) ([]*domain.AdminUser, int64, error) {
	adminUsers, count, err := ia.AdminRepository.GetAllAdminUser(ctx, params)
	if err != nil {
		return nil, 0, err
	}
	return adminUsers, count, nil
}

// IsAdminExist check regist admin
// ia: admin interactor
func (ia *AdminInteractor) IsAdminExist(ctx context.Context, mail string) (bool, error) {
	isExist, err := ia.AdminRepository.IsAdminExist(ctx, mail)
	return isExist, err
}

// PutAdminUser update AdminUser
// ia: admin interactor
func (ia *AdminInteractor) PutAdminUser(ctx context.Context, params domain.AdminUser) (*domain.AdminUser, error) {
	adminUser, err := ia.AdminRepository.PutAdminUser(ctx, params)
	return adminUser, err
}

// DeleteAdmin check regist admin
// ia: admin interactor
func (ia *AdminInteractor) DeleteAdmin(ctx context.Context, uuid string) (*domain.AdminUser, error) {
	adminUser, err := ia.AdminRepository.DeleteAdminUser(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return adminUser, err
}
