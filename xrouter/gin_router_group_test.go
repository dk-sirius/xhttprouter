package xrouter

import (
	"net/http"
	"testing"

	"github.com/dk-sirius/xhttprouter/xhttp"
	"github.com/gin-gonic/gin"
)

type BaseRouterImp struct {
}

func (brmp *BaseRouterImp) Method() string {
	return http.MethodGet
}

func (brmp *BaseRouterImp) Path() string {
	return "/test/base"
}

func (brmp *BaseRouterImp) Context(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pingPong-base")
}

type PathObject struct {
	Path string `path:"/test/basic"`
}

type BasicRouterImp struct {
	xhttp.MethodGet
	Path string `path:"/test/basic"`
}

func (brmp *BasicRouterImp) Context(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pingPong-basic")
}

func TestRouterType(t *testing.T) {
	root := NewRoot(gin.Default())
	base := root.XGroup("/srv-test-base", &BaseRouterImp{})
	base.XHandle(&BaseRouterImp{})

	basic := root.XGroup("/srv-test-basic")
	basic.XHandle(&BasicRouterImp{})
}

