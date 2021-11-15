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
	if ty.Kind() == reflect.Struct {
		for i := 0; i < ty.NumField(); i++ {
			if ty.Field(i).Tag.Get(tag) != "" {
				return ty.Field(i).Tag.Get(tag)
			}
			if ty.Field(i).Type.Kind() == reflect.Ptr || ty.Field(i).Type.Kind() == reflect.Struct {
				traversalTag(ty.Field(i).Type, tag)
			}
		}
	}
	return ""
}
