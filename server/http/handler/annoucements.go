package handler

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func ApplyAnnouncementHandler(r *server.Hertz) {
	// CN
	// Username & Password login (from client).
	r.POST("/hk4e_cn/mdk/shield/api/login", ClientLogin)
	// Cached token login (from registry).
	r.POST("/hk4e_cn/mdk/shield/api/verify", TokenLogin)
	// Combo token login (from session key).
	r.POST("/hk4e_cn/combo/granter/login/v2/login", SessionKeyLogin)
}
