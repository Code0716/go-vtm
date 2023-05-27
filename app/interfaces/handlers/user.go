// // Package handlers hanelers package
package handlers

import (
	"context"
	"fmt"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/registry"
)

type userHandler struct {
	reg registry.Getter
}

func (h userHandler) CreateUser(c context.Context, input domain.User) (*domain.User, error) {
	fmt.Println("ここやで！！！！！！！")

	// isUUID := util.IsValidUUID(uuid)
	// if !isUUID {
	// 	return nil, domain.NewError(domain.ErrorTypeUUIDValidationFailed)
	// }

	userInteractor := h.reg.UserInteractor()
	var user domain.User
	newUser, err := userInteractor.CreateUser(c, user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
