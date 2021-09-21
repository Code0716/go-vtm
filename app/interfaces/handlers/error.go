package handlers

import (
	"errors"
	"net/http"

	"github.com/Code0716/go-vtm/app/domain"
	"github.com/labstack/echo/v4"
)

func sendError(c echo.Context, err error) error {
	var e domain.Error
	if !errors.As(err, &e) {
		return c.JSON(http.StatusInternalServerError, buildInternalServerErrorResponse())
	}

	var code int
	switch e.Type {
	case domain.ErrorTypeValidationFailed:
		fallthrough
	case domain.ErrorTypeAdminEmailValidationFailed:
		fallthrough
	case domain.ErrorTypeRegistItemAlreadyRegistered:
		fallthrough
	case domain.ErrorTypeRegistAdminValidationFailed:
		fallthrough
	case domain.ErrorTypeAdminLoginValidationFailed:
		fallthrough
	case domain.ErrorTypeUUIDValidationFailed:
		fallthrough
	case domain.ErrorTypeNotFound:
		fallthrough
	case domain.ErrorTypePasswordOrEmailValidationFailed:
		code = http.StatusBadRequest
	case domain.ErrorTypeAuthenticationFailed:
		code = http.StatusUnauthorized
	case domain.ErrorTypeInternalError:
		return c.JSON(http.StatusInternalServerError, buildInternalServerErrorResponse())
	default:
		return c.JSON(http.StatusInternalServerError, buildInternalServerErrorResponse())
	}

	errRes := domain.ErrorResponse{
		Error: domain.Error{
			Type:    e.Type,
			Status:  code,
			Message: domain.ErrorMessageMap[e.Type],
		},
	}

	return c.JSON(code, errRes)
}

func buildInternalServerErrorResponse() domain.ErrorResponse {
	errRes := domain.ErrorResponse{
		Error: domain.Error{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
		},
	}

	return errRes
}
