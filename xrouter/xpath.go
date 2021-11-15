package xrouter

import (
	"reflect"
)

const Token = "path"

func Path(s interface{}) string {
	ty := reflect.TypeOf(s)
	return traversalTag(ty, Token)
}

func traversalTag(ty reflect.Type, tag string) string {
	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}
	path := ""
	if ty.Kind() == reflect.Struct {
		for i := 0; i < ty.NumField(); i++ {
			if v, ok := ty.Field(i).Tag.Lookup(tag); ok {
				path = v
				break
			}
			if ty.Field(i).Type.Kind() == reflect.Ptr || ty.Field(i).Type.Kind() == reflect.Struct {
				path = traversalTag(ty.Field(i).Type, tag)
			}
		}
	}
	return path
}
