package interactors

// import (
// 	"context"
// 	"time"

// 	"github.com/Code0716/go-vtm/app/domain"
// 	"github.com/Code0716/go-vtm/app/usecase/repositories"
// 	"github.com/Code0716/go-vtm/app/util"
// )

// // UsersInteractor is user interactor.
// type UsersInteractor struct {
// 	UsersRepository repositories.UsersRepository
// }

// // NewUsers initializes item interactor.
// func NewUsers(
// 	usersRepo repositories.UsersRepository,
// ) *UsersInteractor {
// 	return &UsersInteractor{
// 		UsersRepository: usersRepo,
// 	}
// }

// // UserGetAll returns user list
// // im: users interactor
// func (im *UsersInteractor) UserGetAll(ctx context.Context, params domain.Pager) ([]*domain.User, int64, error) {
// 	userList, count, err := im.UsersRepository.AdminUserGetAll(ctx, params)
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	return userList, count, nil
// }

// // RegistUser regist user
// // im: users interactor
// func (im *UsersInteractor) RegistUser(ctx context.Context, params domain.User) error {
// 	currentTime := time.Now()
// 	params.CreatedAt = currentTime
// 	params.UpdatedAt = currentTime
// 	params.UserId = util.UUIDGenerator()

// 	params.Status = domain.StatusCodeInit.GetWorkStatus()
// 	err := im.UsersRepository.AdminRegistUser(ctx, params)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // UpdateUser update user
// // im: users interactor
// func (im *UsersInteractor) UpdateUser(ctx context.Context, params domain.UpdateUserJSONBody, uuid string) (*domain.User, error) {

// 	oldUser, err := im.GetUserByUUID(ctx, uuid)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if oldUser.DeletedAt != nil {
// 		return nil, domain.WrapError(domain.ErrorTypeUserAlreadyDeleted, err)
// 	}

// 	if params.Name != "" {
// 		oldUser.Name = params.Name
// 	}

// 	if params.PhoneNumber != "" {
// 		oldUser.PhoneNumber = params.PhoneNumber
// 	}

// 	if params.Status != "" {
// 		oldUser.Status = params.Status
// 	}

// 	if params.HourlyPrice != nil {
// 		oldUser.HourlyPrice = params.HourlyPrice
// 	}

// 	oldUser.UpdatedAt = time.Now()

// 	newUser, err := im.UsersRepository.UpdateUser(ctx, *oldUser)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return newUser, nil
// }

// // IsUserExist check regist user
// // im: users interactor
// func (im *UsersInteractor) IsUserExist(ctx context.Context, name, phone string) (bool, error) {
// 	isExist, err := im.UsersRepository.IsUserExist(ctx, name, phone)
// 	return isExist, err
// }

// // GetUserByUUID get regist user by uuid
// // im: users interactor
// func (im *UsersInteractor) GetUserByUUID(ctx context.Context, uuid string) (*domain.User, error) {
// 	user, err := im.UsersRepository.GetUserByUUID(ctx, uuid)
// 	return user, err
// }
