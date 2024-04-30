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
	r = server.Default(server.WithHostPorts(":8080"))

	r.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})
	handler.ApplyAnnouncementHandler(r)
	handler.ApplyGenericHandler(r)
	handler.ApplyLogRoutes(r)
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
	proxy, err := reverseproxy.NewSingleHostReverseProxy(url)
	r.NoRoute(func(c context.Context, ctx *app.RequestContext) {
		fmt.Println(ctx.URI())
		body, _ := ctx.Body()
		fmt.Println(string(body))
		proxy.ServeHTTP(c, ctx)
	})
	if err != nil {
		panic(err)
	}
}
