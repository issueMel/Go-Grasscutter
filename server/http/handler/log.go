package handler

import "github.com/cloudwego/hertz/pkg/app/server"

func ApplyLogRoutes(r *server.Hertz) {
	r.POST("/log", Response("{\"code\":0}"))
	r.POST("/crash/dataUpload", Response("{\"code\":0}"))
}
