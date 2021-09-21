// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// adminUser一覧の取得
	// (GET /admin)
	GetAdminList(ctx echo.Context, params GetAdminListParams) error
	// admin login
	// (POST /admin/login)
	AdminLogin(ctx echo.Context) error
	// member一覧の取得
	// (GET /admin/members)
	AdminGetMemberList(ctx echo.Context, params AdminGetMemberListParams) error
	// memberの登録
	// (POST /admin/members)
	AdminRegistMember(ctx echo.Context) error
	// admin登録
	// (POST /admin/regist)
	RegistAdmin(ctx echo.Context) error
	// adminUserの取得
	// (GET /admin/{uuid})
	GetAdminInfo(ctx echo.Context, uuid string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetAdminList converts echo context to params.
func (w *ServerInterfaceWrapper) GetAdminList(ctx echo.Context) error {
	var err error

	ctx.Set(SecurityScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetAdminListParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "status" -------------

	err = runtime.BindQueryParameter("form", true, false, "status", ctx.QueryParams(), &params.Status)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter status: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAdminList(ctx, params)
	return err
}

// AdminLogin converts echo context to params.
func (w *ServerInterfaceWrapper) AdminLogin(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AdminLogin(ctx)
	return err
}

// AdminGetMemberList converts echo context to params.
func (w *ServerInterfaceWrapper) AdminGetMemberList(ctx echo.Context) error {
	var err error

	ctx.Set(SecurityScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params AdminGetMemberListParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "status" -------------

	err = runtime.BindQueryParameter("form", true, false, "status", ctx.QueryParams(), &params.Status)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter status: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AdminGetMemberList(ctx, params)
	return err
}

// AdminRegistMember converts echo context to params.
func (w *ServerInterfaceWrapper) AdminRegistMember(ctx echo.Context) error {
	var err error

	ctx.Set(SecurityScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AdminRegistMember(ctx)
	return err
}

// RegistAdmin converts echo context to params.
func (w *ServerInterfaceWrapper) RegistAdmin(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RegistAdmin(ctx)
	return err
}

// GetAdminInfo converts echo context to params.
func (w *ServerInterfaceWrapper) GetAdminInfo(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "uuid" -------------
	var uuid string

	err = runtime.BindStyledParameterWithLocation("simple", false, "uuid", runtime.ParamLocationPath, ctx.Param("uuid"), &uuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	ctx.Set(SecurityScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAdminInfo(ctx, uuid)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/admin", wrapper.GetAdminList)
	router.POST(baseURL+"/admin/login", wrapper.AdminLogin)
	router.GET(baseURL+"/admin/members", wrapper.AdminGetMemberList)
	router.POST(baseURL+"/admin/members", wrapper.AdminRegistMember)
	router.POST(baseURL+"/admin/regist", wrapper.RegistAdmin)
	router.GET(baseURL+"/admin/:uuid", wrapper.GetAdminInfo)

}

