package xrouter_test

import (
	"net/http"
	"testing"

	"github.com/dk-sirius/xhttprouter/xhttp"
	"github.com/dk-sirius/xhttprouter/xrouter"
	"github.com/gin-gonic/gin"
)

type BaseRouterImp struct {
	A string
}

func (brmp *BaseRouterImp) Method() string {
	return http.MethodGet
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
	root := xrouter.NewRoot(gin.Default())
	base := root.XGroup("/srv-test-base", &BaseRouterImp{})
	base.XHandle(&BaseRouterImp{})

	basic := root.XGroup("/srv-test-basic")
	basic.XHandle(&BasicRouterImp{})
}
