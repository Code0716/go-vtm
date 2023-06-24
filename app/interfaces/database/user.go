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
func (r *UserRepository) CreateUser(c context.Context, user domain.User) (*domain.User, error) {
	err := r.SQLHandler.Create(c, user).Conn.Error
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var newUser domain.User
	err = r.SQLHandler.First(c, &newUser).Conn.Where("user_id =?", user.UserID).Error
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &newUser, nil
}
