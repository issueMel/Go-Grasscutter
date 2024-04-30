package http

import (
	"Go-Grasscutter/src/server/http/handler"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/reverseproxy"
)

var r *server.Hertz

func InitRouter() *server.Hertz {
	// https://www.cloudwego.io/zh/docs/hertz/reference/config/
	r = server.Default(server.WithHostPorts(":8080"))

	r.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})
	handler.GetAnnouncement(r)
	return r
}

func GetRouter() *server.Hertz {
	if r == nil {
		InitRouter()
	}
	return r
}

func ProxyNoRouteRequest(url string) {
	// initialize a reverse proxy and pass the actual backend server url here
	// 初始化反向代理并传入真正后端服务的地址
	proxy, err := reverseproxy.NewSingleHostReverseProxy(url)
	r.NoRoute(func(c context.Context, ctx *app.RequestContext) {
		fmt.Println(ctx.URI())
		proxy.ServeHTTP(c, ctx)
	})
	if err != nil {
		panic(err)
	}
}
