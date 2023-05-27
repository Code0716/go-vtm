package handlers

import (
	"net/http"

	"github.com/Code0716/go-vtm/app/registry"
)

// Handler bands all handler's implementation.
type Handler struct {
	userHandler
}

// New initializes and returns handlers collection.
func New(reg registry.Getter) Handler {
	h := Handler{}
	// reg のセットを忘れるとnilアクセスで落ちる
	h.userHandler.reg = reg

	return h
}

// Context handler context
// TODO:これに置き換える
type Context interface {
	Param(string) string
	Bind(any) error
	Request() *http.Request
	JSON(int, any) error
}
