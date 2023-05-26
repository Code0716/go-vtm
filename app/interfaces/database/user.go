package database

import (
	"context"
	"log"

	"github.com/Code0716/go-vtm/app/domain"
)

// UserRepository is user database.
type UserRepository struct {
	SQLHandler SQLHandlerInterface
}

// NewUser initializes attendance database.
func NewUser(sqlHandler SQLHandlerInterface) *UserRepository {
	return &UserRepository{
		sqlHandler,
	}
}

// CreateUser  create user
func (r *UserRepository) CreateUser(_ context.Context, user domain.User) (*domain.User, error) {
	err := r.SQLHandler.Create(&user).Conn.Error
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &user, nil
}
