package router

import (
	"Go-Grasscutter/server/http/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

var r *server.Hertz

func InitRouter() *server.Hertz {
	r = server.Default(server.WithHostPorts(":443"))
	// change log level when dev
	hlog.SetLevel(hlog.LevelFatal)
	ApplyHandler()

	return r
}

func ApplyHandler() {
	handler.ApplyAnnouncementHandler(r)
	handler.ApplyGenericHandler(r)
	handler.ApplyLogRoutes(r)
	handler.ApplyRegionHandler(r)
}
