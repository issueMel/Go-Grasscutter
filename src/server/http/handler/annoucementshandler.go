package handler

import (
	"Go-Grasscutter/src/server/http/dispatch"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func GetAnnouncement(r *server.Hertz) {
	// CN
	// Username & Password login (from client).
	auth := r.Group("/hk4e_cn/mdk/shield/api/")
	{
		auth.POST("login", dispatch.ClientLogin)
	}
	r.POST("/sdk/dataUpload", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(200, "{\"code\":0}")
	})
	r.POST("/account/risky/api/check", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(200, "{\"retcode\":0,\"message\":\"OK\",\"data\":{\"id\":\"none\",\"action\":\"ACTION_NONE\",\"geetest\":null}}")
	})
}
