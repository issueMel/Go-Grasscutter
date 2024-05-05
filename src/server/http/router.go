package http

import (
	"Go-Grasscutter/src/server/http/handler"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/reverseproxy"
)

var r *server.Hertz

func Mid(c context.Context, ctx *app.RequestContext) {
	ctx.Next(c)
}

func InitRouter() *server.Hertz {
	r = server.Default(server.WithHostPorts(":8080"))
	// change log level when dev
	hlog.SetLevel(hlog.LevelFatal)
	r.Use(Mid)
	r.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})
	//handler.ApplyAnnouncementHandler(r)
	handler.ApplyGenericHandler(r)
	handler.ApplyLogRoutes(r)
	handler.ApplyRegionHandler(r) // todo fix regionHandler so you can enter the game
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
		proxy.SetModifyResponse(func(response *protocol.Response) error {
			fmt.Println(string(response.Body()))
			return nil
		})
	})
	if err != nil {
		panic(err)
	}
}
