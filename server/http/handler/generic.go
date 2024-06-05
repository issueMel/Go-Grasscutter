package handler

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func ApplyGenericHandler(r *server.Hertz) {
	// hk4e-sdk-os.hoyoverse.com
	r.GET("/hk4e_global/mdk/agreement/api/getAgreementInfos", Response(
		"{\"retcode\":0,\"message\":\"OK\",\"data\":{\"marketing_agreements\":[]}}"))

	// hk4e-sdk-os.hoyoverse.com (this could be either GET or POST based on the observation of
	// different clients)
	r.Any("/hk4e_global/combo/granter/api/compareProtocolVersion", Response(
		"{\"retcode\":0,\"message\":\"OK\",\"data\":{\"modified\":true,\"protocol\":{\"id\":0,\"app_id\":4,\"language\":\"en\",\"user_proto\":\"\",\"priv_proto\":\"\",\"major\":7,\"minimum\":0,\"create_time\":\"0\",\"teenager_proto\":\"\",\"third_proto\":\"\"}}}"))

	// api-account-os.hoyoverse.com
	r.POST("/account/risky/api/check", Response(
		"{\"retcode\":0,\"message\":\"OK\",\"data\":{\"id\":\"none\",\"action\":\"ACTION_NONE\",\"geetest\":null}}"))

	// sdk-os-static.hoyoverse.com
	r.GET("/combo/box/api/config/sdk/combo", Response(
		"{\"retcode\":0,\"message\":\"OK\",\"data\":{\"vals\":{\"disable_email_bind_skip\":\"false\",\"email_bind_remind_interval\":\"7\",\"email_bind_remind\":\"true\"}}}"))

	// hk4e-sdk-os-static.hoyoverse.com
	r.GET("/hk4e_global/mdk/shield/api/getConfig", Response(
		"{\"retcode\":0,\"message\":\"OK\",\"data\":{\"protocol\":true,\"qr_enabled\":false,\"log_level\":\"INFO\",\"announce_url\":\"https://webstatic-sea.hoyoverse.com/hk4e/announcement/index.html?sdk_presentation_style=fullscreen\\u0026sdk_screen_transparent=true\\u0026game_biz=hk4e_global\\u0026auth_appid=announcement\\u0026game=hk4e#/\",\"push_alias_type\":2,\"disable_ysdk_guard\":false,\"enable_announce_pic_popup\":true}}"))

	// hk4e-sdk-os-static.hoyoverse.com
	r.GET("/hk4e_global/mdk/shield/api/loadConfig", Response(
		"{\"retcode\":0,\"message\":\"OK\",\"data\":{\"id\":6,\"game_key\":\"hk4e_global\",\"client\":\"PC\",\"identity\":\"I_IDENTITY\",\"guest\":false,\"ignore_versions\":\"\",\"scene\":\"S_NORMAL\",\"name\":\"原神海外\",\"disable_regist\":false,\"enable_email_captcha\":false,\"thirdparty\":[\"fb\",\"tw\"],\"disable_mmt\":false,\"server_guest\":false,\"thirdparty_ignore\":{\"tw\":\"\",\"fb\":\"\"},\"enable_ps_bind_account\":false,\"thirdparty_login_configs\":{\"tw\":{\"token_type\":\"TK_GAME_TOKEN\",\"game_token_expires_in\":604800},\"fb\":{\"token_type\":\"TK_GAME_TOKEN\",\"game_token_expires_in\":604800}}}}"))

	// Upload device info
	// abtest-api-data-sg.hoyoverse.com
	r.POST("/data_abtest_api/config/experiment/list", Response(
		"{\"retcode\":0,\"success\":true,\"message\":\"\",\"data\":[{\"code\":1000,\"type\":2,\"config_id\":\"14\",\"period_id\":\"6036_99\",\"version\":\"1\",\"configs\":{\"cardType\":\"old\"}}]}"))

	// log-upload-os.mihoyo.com
	r.Any("/log/sdk/upload", Response("{\"code\":0}"))
	r.Any("/sdk/upload", Response("{\"code\":0}"))
	r.POST("/sdk/dataUpload", Response("{\"code\":0}"))

	// /perf/config/verify?device_id=xxx&platform=x&name=xxx
	r.Any("/perf/config/verify", Response("{\"code\":0}"))

	// todo INCOMPLETE: webstatic-sea.hoyoverse.com
	//javalin.get("/admin/mi18n/plat_oversea/*", new WebStaticVersionResponse());
	//javalin.get("/status/server", GenericHandler::serverStatus);
}
