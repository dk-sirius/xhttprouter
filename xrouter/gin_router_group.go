package xrouter

import (
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	group *gin.RouterGroup
}

func NewRoot(ctx *gin.Engine) *RouterGroup {
	return &RouterGroup{
		group: ctx.Group(""),
	}
}

func (rg *RouterGroup) XGroup(relativePath string, handlers ...interface{}) *RouterGroup {
	tmpGroup := rg.group.Group(relativePath)
	if len(handlers) > 0 {
		for _, router := range handlers {
			if rr, ok := router.(XContext); ok {
				tmpGroup.Handlers = append(tmpGroup.Handlers, rr.Context)
			} else {
				panic("invalid router")
			}
		}
	}
	return &RouterGroup{
		group: tmpGroup,
	}
}

func (rg *RouterGroup) XHandle(r interface{}) {
	// context，http method
	basicRouter, err := validRoute(r)
	if err != nil {
		panic("invalid route")
	}
	// request path
	relativePath := getPath(r)
	rg.group.Handle(basicRouter.Method(), relativePath, basicRouter.Context)
}
