package handler

import (
	"Go-Grasscutter/src/server/http/dispatch"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func ApplyAnnouncementHandler(r *server.Hertz) {
	// CN
	// Username & Password login (from client).
	r.POST("/hk4e_cn/mdk/shield/api/login", dispatch.ClientLogin)
	// Cached token login (from registry).
	r.POST("/hk4e_cn/mdk/shield/api/verify", dispatch.TokenLogin)
	// Combo token login (from session key).
	r.POST("/hk4e_cn/combo/granter/login/v2/login", dispatch.SessionKeyLogin)
}
