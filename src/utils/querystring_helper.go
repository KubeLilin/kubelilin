package utils

import (
	"github.com/yoyofx/yoyogo/utils/cast"
	"github.com/yoyofx/yoyogo/web/context"
)

func GetNumberOfParam[N cast.Number](ctx *context.HttpContext, name string) N {
	queryStringValue := ctx.Input.QueryDefault(name, "0")
	number, _ := cast.Str2Number[N](queryStringValue)
	return number
}
