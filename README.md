// 采用实现BaseRouter 方式
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

// 采用Router方式「内嵌结构体（已经实现Method），结构体tag（`path:`）,路径可以嵌套...」
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

