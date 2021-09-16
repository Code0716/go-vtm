package interactor

import (
	"context"
	"time"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/interfaces/repository"
	"github.com/Code0716/go-vtm/app/util"
	uuid "github.com/satori/go.uuid"
)

// AdminInteractor is admin interactor.
type AdminInteractor struct {
	AdminRepository repository.AdminInterface
}

// NewAdmin initializes item interactor.
func NewAdmin(
	adminsRepo repository.AdminInterface,
) *AdminInteractor {
	return &AdminInteractor{
		AdminRepository: adminsRepo,
	}
}

// GetAdminJwtByEmail returns jwt
// ia: admin interactor
func (ia *AdminInteractor) GetAdminJwtByEmail(ctx context.Context, params domain.AdminLoginJSONRequestBody) (*string, error) {
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
		return nil, domain.WrapInternalError(err)
	}

	return &token, nil
}

// RegistAdmin returns member list
// ia: admin interactor
func (ia *AdminInteractor) RegistAdmin(ctx context.Context, params domain.RegistAdminJSONRequestBody) error {
	hash, err := util.GetHush(params.Password)
	if err != nil {
		return domain.WrapInternalError(err)
	}
	currentTime := time.Now()

	registAdmin := domain.AdminUser{
		Name:        params.Name,
		Authority:   "admin",
		AdminId:     uuid.NewV4().String(),
		MailAddress: params.MailAddress,
		Status:      "init",
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
func (ia *AdminInteractor) IsAdminExist(ctx context.Context, name, mail string) (bool, error) {
	isExist, err := ia.AdminRepository.IsAdminExist(ctx, name, mail)
	return isExist, err
}
