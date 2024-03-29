// Package handlers hanelers package
package handlers

// import (
// 	"net/http"

// 	"github.com/Code0716/go-vtm/app/constants"
// 	"github.com/Code0716/go-vtm/app/domain"
// 	"github.com/Code0716/go-vtm/app/registry"
// 	"github.com/Code0716/go-vtm/app/util"
// 	"github.com/labstack/echo/v4"
// )

// type adminHandler struct {
// 	reg registry.Getter
// }

// // RegistAdmin is regest admin user handler
// // TODO:二段階認証にしたい。
// func (h adminHandler) RegistAdmin(c echo.Context) error {
// 	var newAdmin domain.RegistAdminJSONRequestBody
// 	err := c.Bind(&newAdmin)
// 	if err != nil {
// 		return sendError(c, domain.NewError(domain.ErrorTypeValidationFailed))
// 	}

// 	if newAdmin.Name == "" || newAdmin.Password == "" || newAdmin.MailAddress == "" {
// 		return sendError(c, domain.NewError(domain.ErrorTypeRegistAdminValidationFailed))
// 	}

// 	if !util.ValidEmailAddress(newAdmin.MailAddress) {
// 		return sendError(c, domain.NewError(domain.ErrorTypeAdminEmailValidationFailed))
// 	}

// 	adminInteractor := h.reg.AdminInteractor()
// 	if isRegested, err := adminInteractor.IsAdminExist(c.Request().Context(), newAdmin.MailAddress); isRegested {
// 		if err != nil {
// 			return sendError(c, err)
// 		}
// 		if isRegested {
// 			return sendError(c, domain.NewError(domain.ErrorTypeRegistItemAlreadyRegistered))
// 		}
// 	}

// 	err = adminInteractor.RegistAdmin(c.Request().Context(), newAdmin)
// 	if err != nil {
// 		return sendError(c, err)
// 	}

// 	return c.JSON(http.StatusCreated, domain.CommonSuccessResponse{Message: constants.RegistSuccess})

// }

// func (h adminHandler) GetAdminList(c echo.Context, params api.GetAdminListParams) error {
// 	var limit int
// 	if params.Limit != nil {
// 		limit = *params.Limit
// 	}

// 	if limit <= 0 || 50 < limit {
// 		limit = 50
// 	}
// 	var offset int
// 	if params.Offset != nil {
// 		offset = *params.Offset
// 	}
// 	var status string
// 	if params.Status != nil {
// 		status = *params.Status
// 	}

// 	aminListParams := domain.Pager{
// 		Limit:  limit,
// 		Offset: offset,
// 		Status: status,
// 	}
// 	adminInteractor := h.reg.AdminInteractor()
// 	al, count, err := adminInteractor.GetAdminList(c.Request().Context(), aminListParams)
// 	if err != nil {
// 		return sendError(c, err)
// 	}

// 	adminUserList := make([]*domain.AdminUser, 0, len(al))
// 	for _, item := range al {
// 		user := domain.AdminUser{
// 			Id:          item.Id,
// 			AdminId:     item.AdminId,
// 			Name:        item.Name,
// 			MailAddress: item.MailAddress,
// 			Permission:  item.Permission,
// 			Status:      item.Status,
// 			CreatedAt:   item.CreatedAt,
// 			UpdatedAt:   item.UpdatedAt,
// 			DeletedAt:   item.DeletedAt,
// 		}
// 		adminUserList = append(adminUserList, &user)
// 	}
// 	res := domain.AdminUsersResponse{
// 		AdminUsers: adminUserList,
// 		Total:      count,
// 	}

// 	return c.JSON(http.StatusOK, res)
// }

// func (h adminHandler) GetAdminUser(c echo.Context, uuid string) error {
// 	isUUID := util.IsValidUUID(uuid)
// 	if !isUUID {
// 		return sendError(c, domain.NewError(domain.ErrorTypeUUIDValidationFailed))
// 	}

// 	adminInteractor := h.reg.AdminInteractor()
// 	adminUser, err := adminInteractor.GetAdminByUUID(c.Request().Context(), uuid)
// 	if err != nil {
// 		return sendError(c, err)
// 	}

// 	res := domain.AdminUserResponse{
// 		Id:          adminUser.Id,
// 		Name:        adminUser.Name,
// 		AdminId:     adminUser.AdminId,
// 		MailAddress: adminUser.MailAddress,
// 		Status:      adminUser.Status,
// 		Permission:  adminUser.Permission,
// 		CreatedAt:   adminUser.CreatedAt,
// 		UpdatedAt:   adminUser.UpdatedAt,
// 		DeletedAt:   adminUser.DeletedAt,
// 	}

// 	return c.JSON(http.StatusOK, res)
// }

// // UpdateAdminUser is update adminUser
// func (h adminHandler) UpdateAdminUser(c echo.Context, uuid string) error {
// 	isUUID := util.IsValidUUID(uuid)
// 	if !isUUID {
// 		return sendError(c, domain.NewError(domain.ErrorTypeUUIDValidationFailed))
// 	}

// 	adminInteractor := h.reg.AdminInteractor()
// 	adminUser, err := adminInteractor.GetAdminByUUID(c.Request().Context(), uuid)
// 	if err != nil {
// 		return sendError(c, err)
// 	}

// 	var updateAdmin domain.UpdateAdminUserJSONRequestBody
// 	err = c.Bind(&updateAdmin)
// 	if err != nil {
// 		return sendError(c, domain.NewError(domain.ErrorTypeValidationFailed))
// 	}

// 	adminUser.Name = updateAdmin.Name
// 	if !util.ValidEmailAddress(updateAdmin.MailAddress) {
// 		return sendError(c, domain.NewError(domain.ErrorTypeAdminEmailValidationFailed))
// 	}
// 	adminUser.MailAddress = updateAdmin.MailAddress
// 	adminUser.Permission = updateAdmin.Permission
// 	adminUser.Status = updateAdmin.Status

// 	response, err := adminInteractor.PutAdminUser(c.Request().Context(), *adminUser)
// 	if err != nil {
// 		return sendError(c, err)
// 	}

// 	return c.JSON(http.StatusOK, response)
// }

// // DeleteAdminUser delete admin info
// func (h adminHandler) DeleteAdminUser(c echo.Context, uuid string) error {

// 	isUUID := util.IsValidUUID(uuid)
// 	if !isUUID {
// 		return sendError(c, domain.NewError(domain.ErrorTypeUUIDValidationFailed))
// 	}
// 	adminInteractor := h.reg.AdminInteractor()

// 	adminUser, err := adminInteractor.DeleteAdmin(c.Request().Context(), uuid)
// 	if err != nil {
// 		return sendError(c, err)
// 	}

// 	response := domain.DeleteAdminUserResponse{
// 		Id:      adminUser.Id,
// 		AdminId: adminUser.AdminId,
// 		Name:    adminUser.Name,
// 	}

// 	return c.JSON(http.StatusOK, response)
// }

// // func (h adminHandler) Login(c echo.Context) error {
// // 	var loginRequest domain.LoginJSONRequestBody
// // 	err := c.Bind(&loginRequest)
// // 	if err != nil {
// // 		return sendError(c, domain.NewError(domain.ErrorTypeValidationFailed))
// // 	}

// // 	if loginRequest.MailAddress == "" || loginRequest.Password == "" {
// // 		return sendError(c, domain.NewError(domain.ErrorTypeLoginValidationFailed))
// // 	}

// // 	adminInteractor := h.reg.AdminInteractor()
// // 	token, err := adminInteractor.GetAdminJwtByEmail(c.Request().Context(), loginRequest)
// // 	if err != nil {
// // 		return sendError(c, err)
// // 	}

// // 	reseponse := domain.AuthenticationResponse{
// // 		Token:   token,
// // 		Message: "Success",
// // 	}

// // 	return c.JSON(http.StatusOK, reseponse)
// // }
