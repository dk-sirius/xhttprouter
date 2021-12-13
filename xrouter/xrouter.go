package xrouter

import (
	"fmt"
)

type PathRouter interface {
	Path() string
	MethodRouter
}

type XContext interface {
	Context(ctx *Context)
}

type MethodRouter interface {
	XContext
	Method() string
}

func validRoute(r interface{}) (MethodRouter, error) {
	if route, ok := r.(MethodRouter); ok {
		return route, nil
	}
	return nil, fmt.Errorf("invalid route")
}

func getPath(r interface{}) string {
	switch rr := r.(type) {
	case PathRouter:
		return rr.Path()
	case MethodRouter:
		return Path(r)
	default:
		return ""
	}
}
