package dispatch

import (
	"Go-Grasscutter/src/auth"
	"Go-Grasscutter/src/server/http/object"
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"net/http"
)

func ClientLogin(c context.Context, ctx *app.RequestContext) {
	var req *object.LoginAccountRequestJson
	// Validate data.
	err := ctx.BindJSON(&req)
	if err != nil {
		log.Println(err)
		return
	}
	// Pass data to authentication handler.
	var pa = new(auth.PasswordAuthenticator)
	resp := pa.Authenticate(&auth.AuthenticationRequest{
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
		log.Println(err)
		return
	}
	// Decode additional body data.
	tokenData := &object.LoginTokenData{}
	err = json.Unmarshal([]byte(req.Data), &tokenData)
	if err != nil {
		log.Println(err)
		return
	}
	// Pass data to authentication handler.
	var ska = new(auth.SessionKeyAuthenticator)
	resp := ska.Authenticate(&auth.AuthenticationRequest{
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
		log.Println(err)
		return
	}
	// Pass data to authentication handler.
	var ta = new(auth.TokenAuthenticator)
	resp := ta.Authenticate(&auth.AuthenticationRequest{
		TokenRequest: req,
	}).(*object.LoginResultJson)
	ctx.JSON(200, resp)
}
