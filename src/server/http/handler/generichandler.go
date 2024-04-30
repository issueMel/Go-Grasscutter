package handler

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func ApplyGenericHandler(r *server.Hertz) {
	// api-account-os.hoyoverse.com
	r.POST("/account/risky/api/check", Response("{\"retcode\":0,\"message\":\"OK\",\"data\":{\"id\":\"none\",\"action\":\"ACTION_NONE\",\"geetest\":null}}"))
	// Test api?
	// abtest-api-data-sg.hoyoverse.com
	r.POST("/data_abtest_api/config/experiment/list", Response("{\"retcode\":0,\"success\":true,\"message\":\"\",\"data\":[{\"code\":1000,\"type\":2,\"config_id\":\"14\",\"period_id\":\"6036_99\",\"version\":\"1\",\"configs\":{\"cardType\":\"old\"}}]}"))
	// log-upload-os.mihoyo.com
	r.POST("/sdk/dataUpload", Response("{\"code\":0}"))
}
