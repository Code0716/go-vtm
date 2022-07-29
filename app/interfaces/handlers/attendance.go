package handlers

import (
	"log"
	"net/http"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/Code0716/go-vtm/app/registry"
	"github.com/labstack/echo/v4"
)

type attendanceHandler struct {
	reg registry.Getter
}

func (h attendanceHandler) Timestamp(c echo.Context) error {
	var timestamp domain.TimestampJSONBody

	err := c.Bind(&timestamp)
	if err != nil {
		log.Print(err)
		return sendError(c, domain.WrapInternalError(err))
	}

	if timestamp.MemberId == "" || timestamp.Status == "" {
		log.Print("Timestamp params not set")
		return sendError(c, domain.NewError(domain.ErrorTypeRequestParamsNotSet))
	}

	// TODO: 後ほど修正する。
	if timestamp.Status != "BEGIN_WORK" &&
		timestamp.Status != "BEGIN_REST" &&
		timestamp.Status != "END_WORK" &&
		timestamp.Status != "END_REST" {
		return sendError(c, domain.NewError(domain.ErrorTypeRequestParamsNotSet))
	}

	attendanceInteractor := h.reg.AttendanceInteractor()
	attendance, err := attendanceInteractor.Timestamp(c.Request().Context(), timestamp)

	if err != nil {
		log.Print(err)
		return sendError(c, domain.WrapInternalError(err))
	}
	return c.JSON(http.StatusOK, attendance)
}
