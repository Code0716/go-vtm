// // Package handlers hanelers package
package handlers

import (
	"context"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/registry"
	"github.com/Code0716/go-vtm/app/util"
)

type userHandler struct {
	reg registry.Getter
}

func (h userHandler) GetAdminUser(c context.Context, uuid string) (*domain.User, error) {
	isUUID := util.IsValidUUID(uuid)
	if !isUUID {
		return nil, domain.NewError(domain.ErrorTypeUUIDValidationFailed)
	}

	userInteractor := h.reg.UserInteractor()
	var user domain.User
	newUser, err := userInteractor.CreateUser(c, user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
