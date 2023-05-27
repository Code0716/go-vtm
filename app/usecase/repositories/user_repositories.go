// Package repositories db interfaces
package repositories

import (
	"context"

	"github.com/Code0716/go-vtm/app/domain"
)

// UserRepository  is data access methods to user.
type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) (*domain.User, error)
}
