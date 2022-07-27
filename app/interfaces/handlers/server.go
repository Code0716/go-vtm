package handlers

import (
	"net/http"

	"github.com/Code0716/go-vtm/app/registry"
)

// Handler bands all handler's implementation.
type Handler struct {
	membersHandler
	adminHandler
	loginHandler
	attendanceHandler
}

// New initializes and returns handlers collection.
func New(reg registry.Getter) Handler {
	h := Handler{}
	// reg のセットを忘れるとnilアクセスで落ちる
	h.adminHandler.reg = reg
	h.membersHandler.reg = reg
	h.loginHandler.reg = reg
	h.adminHandler.reg = reg

	return h
}

// Context handler context
// TODO:これに置き換える
type Context interface {
	Param(string) string
	Bind(interface{}) error
	Request() *http.Request
	JSON(int, interface{}) error
}
