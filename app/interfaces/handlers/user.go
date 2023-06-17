// Package handlers hanelers package
package handlers

import (
	"context"
	"time"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/registry"
	"github.com/Code0716/go-vtm/app/util"
	"github.com/Code0716/go-vtm/graph/model"
)

type userHandler struct {
	reg registry.Getter
}

// CreateUser　handler
func (h userHandler) CreateUser(c context.Context, input model.CreateUserInput) (*model.User, error) {

	user := domain.User{
		UserID:       util.UUIDGenerator(),
		Name:         input.Name,
		MailAddress:  input.MailAddress,
		PhoneNumber:  input.PhoneNumber,
		UnitPrice:    input.UnitPrice,
		DepartmentID: input.DepartmentID,
	}

	if input.Status == nil {
		user.Status = domain.UserStatusInit
	}

	if input.Role == nil {
		user.Role = domain.UserRoleCommon
	}

	if input.EmploymentStatus == nil {
		user.EmploymentStatus = domain.EmploymentStatusHourly
	}

	currentTime := time.Now()
	user.CreatedAt = currentTime
	user.UpdatedAt = currentTime

	userInteractor := h.reg.UserInteractor()
	newUser, err := userInteractor.CreateUser(c, user)
	if err != nil {
		return nil, err
	}

	createdUser := &model.User{
		ID:               *newUser.ID,
		UserID:           newUser.UserID,
		Name:             newUser.Name,
		MailAddress:      newUser.MailAddress,
		PhoneNumber:      newUser.PhoneNumber,
		Status:           model.UserStatus(newUser.Status),
		Role:             model.UserRole(newUser.Role),
		EmploymentStatus: model.EmploymentStatus(newUser.EmploymentStatus),
		UnitPrice:        newUser.UnitPrice,
		DepartmentID:     newUser.DepartmentID,
		CreatedAt:        util.TimeToString(newUser.CreatedAt),
		UpdatedAt:        util.TimeToString(newUser.UpdatedAt),
	}

	if newUser.DeletedAt != nil {
		s := util.TimeToString(*newUser.DeletedAt)
		createdUser.DeletedAt = &s
	}

	return createdUser, nil
}
