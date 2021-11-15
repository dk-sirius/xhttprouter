package xrouter

import (
	"fmt"
)

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

func basicRoute(r interface{}) (Router, error) {
	if route, ok := r.(Router); ok {
		return route, nil
	}
	return nil, fmt.Errorf("invalid route")
}

func getPath(r interface{}) (string, error) {
	if baseRouter, ok := r.(BaseRouter); ok {
		return baseRouter.Path(), nil
	} else if _, ok := r.(Router); ok {
		if p := Path(r); p != "" {
			return p, nil
		}
	}
	return "", fmt.Errorf("not found route path")
}
