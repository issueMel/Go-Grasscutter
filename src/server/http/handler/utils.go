package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func Response(json string) func(c context.Context, ctx *app.RequestContext) {
	return func(c context.Context, ctx *app.RequestContext) {
		ctx.String(200, json)
	}
}
