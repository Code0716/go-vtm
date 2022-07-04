package handlers

// type loginHandler struct {
// 	reg registry.Getter
// }

// func (h loginHandler) Login(c echo.Context) error {
// 	var loginRequest domain.LoginJSONRequestBody
// 	err := c.Bind(&loginRequest)
// 	if err != nil {
// 		return sendError(c, domain.NewError(domain.ErrorTypeValidationFailed))
// 	}

// 	if loginRequest.MailAddress == "" || loginRequest.Password == "" {
// 		return sendError(c, domain.NewError(domain.ErrorTypeLoginValidationFailed))
// 	}

// 	adminInteractor := h.reg.AdminInteractor()
// 	token, err := adminInteractor.GetAdminJwtByEmail(c.Request().Context(), loginRequest)
// 	if err != nil {
// 		return sendError(c, err)
// 	}

// 	reseponse := domain.AuthenticationResponse{
// 		Token:   token,
// 		Message: "Success",
// 	}

// 	return c.JSON(http.StatusOK, reseponse)
// }
