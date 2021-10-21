package handlers

import (
	"net/http"

	"github.com/Code0716/go-vtm/app/constants"
	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/gen/api"
	"github.com/Code0716/go-vtm/app/registry"
	"github.com/labstack/echo/v4"
)

type membersHandler struct {
	reg registry.Getter
}

func (h membersHandler) AdminGetMemberList(c echo.Context, params api.AdminGetMemberListParams) error {
	var limit int
	if params.Limit != nil {
		limit = int(*params.Limit)
	}

	if limit <= 0 || 50 < limit {
		limit = 50
	}

	var offset int
	if params.Offset != nil {
		offset = int(*params.Offset)
	}

	var status string
	if params.Status != nil {
		status = string(*params.Status)
	}

	memberParams := domain.Pager{
		Limit:  limit,
		Offset: offset,
		Status: status,
	}

	membersInteractor := h.reg.MembersInteractor()
	ml, count, err := membersInteractor.MemberGetAll(c.Request().Context(), memberParams)
	if err != nil {
		return sendError(c, err)
	}

	membersList := make([]*domain.Member, 0, len(ml))
	for _, item := range ml {
		member := domain.Member{
			Id:          item.Id,
			MemberId:    item.MemberId,
			Name:        item.Name,
			PhoneNumber: item.PhoneNumber,
			Status:      item.Status,
			HourlyPrice: item.HourlyPrice,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			DeletedAt:   item.DeletedAt,
		}
		membersList = append(membersList, &member)
	}
	res := domain.MembersResponse{
		Members: membersList,
		Total:   count,
	}

	return c.JSON(http.StatusOK, res)

}

func (h membersHandler) AdminRegistMember(c echo.Context) error {
	var newMember domain.Member
	err := c.Bind(&newMember)
	if err != nil {
		return sendError(c, domain.NewError(domain.ErrorTypeValidationFailed))
	}

	// TODO:電話番号のvalidateも入れたい
	if newMember.Name == "" || newMember.PhoneNumber == "" {
		return sendError(c, domain.NewError(domain.ErrorTypeRegistMemberValidationFailed))
	}

	membersInteractor := h.reg.MembersInteractor()

	isExist, err := membersInteractor.IsMemberExist(c.Request().Context(), newMember.Name, newMember.PhoneNumber)
	if err != nil {
		return sendError(c, domain.WrapInternalError(err))
	}

	if isExist {
		return sendError(c, domain.NewError(domain.ErrorTypeRegistItemAlreadyRegistered))
	}

	err = membersInteractor.RegistMember(c.Request().Context(), newMember)
	if err != nil {
		return sendError(c, err)
	}

	res := domain.CommonSuccessResponse{
		Message: constants.REGIST_SUCCESS,
	}

	return c.JSON(http.StatusCreated, res)

}
