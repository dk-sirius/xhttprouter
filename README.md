xhttprouter

当前主要是基于gin-gonic，对于原有的HandleFunc以及RouterGroup做拆分以及封装，将router请求的method，relativePath以及context，封包在接口中:
RouterContext=>Router=>BaseRouter

    //./xrouter/xrouter.go
    type BaseRouter interface {
    	Path() string
    	Router
    }
    
    type RouterContext interface {
    	Context(ctx *Context)
    }
    
    type Router interface {
    	RouterContext
    	Method() string
    }
    // ……

使用方式

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
    
