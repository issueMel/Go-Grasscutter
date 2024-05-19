package handler

import (
	"Go-Grasscutter/log"
	"Go-Grasscutter/server/http/object"
	"Go-Grasscutter/server/http/service"
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func ClientLogin(c context.Context, ctx *app.RequestContext) {
	var req *object.LoginAccountRequestJson
	// Validate data.
	err := ctx.BindJSON(&req)
	if err != nil {
		log.SugaredLogger.Error(err)
		return
	}
	// Pass data to authentication handler.
	var pa = new(service.PasswordAuthenticator)
	resp := pa.Authenticate(&service.AuthenticationRequest{
		Context:         ctx,
		PasswordRequest: req,
	}).(*object.LoginResultJson)
	ctx.JSON(http.StatusOK, resp)
}

func SessionKeyLogin(c context.Context, ctx *app.RequestContext) {
	var req *object.ComboTokenReqJson
	// Validate body data.
	err := ctx.BindJSON(&req)
	if err != nil {
		log.SugaredLogger.Error(err)
		return
	}
	// Decode additional body data.
	tokenData := &object.LoginTokenData{}
	err = json.Unmarshal([]byte(req.Data), &tokenData)
	if err != nil {
		log.SugaredLogger.Error(err)
		return
	}
	// Pass data to authentication handler.
	var ska = new(service.SessionKeyAuthenticator)
	resp := ska.Authenticate(&service.AuthenticationRequest{
		Context:           ctx,
		SessionKeyRequest: req,
		SessionKeyData:    tokenData,
	}).(*object.ComboTokenResJson)
	ctx.JSON(200, &resp)
}

func TokenLogin(c context.Context, ctx *app.RequestContext) {
	var req *object.LoginTokenRequestJson
	// Validate body data.
	err := ctx.BindJSON(&req)
	if err != nil {
		log.SugaredLogger.Error(err)
		return
	}
	// Pass data to authentication handler.
	var ta = new(service.TokenAuthenticator)
	resp := ta.Authenticate(&service.AuthenticationRequest{
		TokenRequest: req,
	}).(*object.LoginResultJson)
	ctx.JSON(200, resp)
}
