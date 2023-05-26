package handlers

// import (
// 	"net/http"

// 	"github.com/Code0716/go-vtm/app/constants"
// 	"github.com/Code0716/go-vtm/app/domain"
// 	"github.com/Code0716/go-vtm/app/registry"
// 	"github.com/Code0716/go-vtm/app/util"
// 	"github.com/labstack/echo/v4"
// )

// type usersHandler struct {
// 	reg registry.Getter
// }

// func (h usersHandler) AdminGetUserList(c echo.Context, params api.AdminGetUserListParams) error {
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

// 	userParams := domain.Pager{
// 		Limit:  limit,
// 		Offset: offset,
// 		Status: status,
// 	}

// 	usersInteractor := h.reg.UsersInteractor()
// 	ml, count, err := usersInteractor.UserGetAll(c.Request().Context(), userParams)
// 	if err != nil {
// 		return sendError(c, err)
// 	}

// 	usersList := make([]*domain.User, 0, len(ml))
// 	for _, item := range ml {
// 		user := domain.User{
// 			Id:          item.Id,
// 			UserId:    item.UserId,
// 			Name:        item.Name,
// 			PhoneNumber: item.PhoneNumber,
// 			Status:      item.Status,
// 			HourlyPrice: item.HourlyPrice,
// 			CreatedAt:   item.CreatedAt,
// 			UpdatedAt:   item.UpdatedAt,
// 			DeletedAt:   item.DeletedAt,
// 		}
// 		usersList = append(usersList, &user)
// 	}
// 	res := domain.UsersResponse{
// 		Users: usersList,
// 		Total:   count,
// 	}

// 	return c.JSON(http.StatusOK, res)

// }

// func (h usersHandler) AdminRegistUser(c echo.Context) error {
// 	var newUser domain.User
// 	err := c.Bind(&newUser)
// 	if err != nil {
// 		return sendError(c, domain.NewError(domain.ErrorTypeValidationFailed))
// 	}

// 	// TODO:電話番号のvalidateも入れたい
// 	if newUser.Name == "" || newUser.PhoneNumber == "" {
// 		return sendError(c, domain.NewError(domain.ErrorTypeRegistUserValidationFailed))
// 	}

// 	usersInteractor := h.reg.UsersInteractor()

// 	isExist, err := usersInteractor.IsUserExist(c.Request().Context(), newUser.Name, newUser.PhoneNumber)
// 	if err != nil {
// 		return sendError(c, domain.NewError(domain.ErrorTypeContentNotFound))
// 	}

// 	if isExist {
// 		return sendError(c, domain.NewError(domain.ErrorTypeRegistItemAlreadyRegistered))
// 	}

// 	err = usersInteractor.RegistUser(c.Request().Context(), newUser)
// 	if err != nil {
// 		return sendError(c, err)
// 	}

// 	res := domain.CommonSuccessResponse{
// 		Message: constants.RegistSuccess,
// 	}

// 	return c.JSON(http.StatusCreated, res)

// }

// func (h usersHandler) UpdateUser(c echo.Context, uuid string) error {
// 	var updateUser domain.UpdateUserJSONBody
// 	err := c.Bind(&updateUser)
// 	if err != nil {
// 		return sendError(c, domain.NewError(domain.ErrorTypeValidationFailed))
// 	}

// 	if !util.IsValidUUID(uuid) {
// 		return sendError(c, domain.NewError(domain.ErrorTypeUUIDValidationFailed))
// 	}

// 	usersInteractor := h.reg.UsersInteractor()
// 	newUser, err := usersInteractor.UpdateUser(c.Request().Context(), updateUser, uuid)
// 	if err != nil {
// 		return sendError(c, err)
// 	}

// 	response := domain.UserResponse{
// 		HourlyPrice: newUser.HourlyPrice,
// 		Id:          newUser.Id,
// 		UserId:    newUser.UserId,
// 		Name:        newUser.Name,
// 		PhoneNumber: newUser.Password,
// 		Status:      newUser.Status,
// 		UpdatedAt:   newUser.UpdatedAt,
// 	}

// 	return c.JSON(http.StatusOK, response)

// }

// func (h usersHandler) GetUser(c echo.Context, uuid string) error {

// 	if !util.IsValidUUID(uuid) {
// 		return sendError(c, domain.NewError(domain.ErrorTypeUUIDValidationFailed))
// 	}
// 	usersInteractor := h.reg.UsersInteractor()
// 	user, err := usersInteractor.GetUserByUUID(c.Request().Context(), uuid)
// 	if err != nil {
// 		return sendError(c, err)
// 	}

// 	return c.JSON(http.StatusOK, user)

// }
