package interactors

import (
	"context"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/usecase/repositories"
)

// UserInteractor is user interactor.
type UserInteractor struct {
	UserRepository repositories.UserRepository
}

// NewUser initializes item interactor.
func NewUser(
	usersRepo repositories.UserRepository,
) *UserInteractor {
	return &UserInteractor{
		UserRepository: usersRepo,
	}
}

// CreateUser returns user list
// iu: users interactor
func (iu *UserInteractor) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	newUser, err := iu.UserRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
