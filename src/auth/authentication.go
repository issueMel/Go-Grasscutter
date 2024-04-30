package auth

import (
	"Go-Grasscutter/src/game"
	"Go-Grasscutter/src/server/http/object"
	"Go-Grasscutter/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
)

type AuthenticationRequest struct {
	Context           *app.RequestContext
	PasswordRequest   *object.LoginAccountRequestJson
	TokenRequest      *object.LoginTokenRequestJson
	SessionKeyRequest *object.ComboTokenReqJson
	SessionKeyData    *object.LoginTokenData
}

// PasswordAuthenticator maybe never use
type PasswordAuthenticator struct {
}

func (PasswordAuthenticator) Authenticate(request *AuthenticationRequest) any {
	// todo redo
	resp := object.NewLoginResultJson()
	req := request.PasswordRequest
	if req == nil {
		log.Println("requestData should never be nil")
		return nil
	}
	var successfulLogin bool
	address := utils.Address(request.Context)
	//responseMessage := translate("messages.dispatch.account.username_error")
	//loggerMessage := ""

	// Get account from db.
	account := game.GetAccountName(req.Account)
	// todo autoCreate
	successfulLogin = account != nil
	if !successfulLogin {
		log.Println("user not exit, req come from ", address)
	}
	if successfulLogin {
		resp.Message = "OK"
		resp.Data.Account.UID = account.ID
		resp.Data.Account.Token = account.GenerateSessionKey()
		resp.Data.Account.Email = account.GetEmail()
	} else {
		resp.RetCode = -201
		resp.Message = "Go-G错误"
	}
	return resp
}

type ExperimentalPasswordAuthenticator struct {
}

//func (a *ExperimentalPasswordAuthenticator) Authenticate(request *AuthenticationRequest) any {
//	// todo 构造后传进来的
//	var req *object.LoginAccountRequestJson
//	req := request.PasswordRequest
//	if req != nil {
//		log.Fatal("requestData should never be nil")
//		return nil
//	}
//	// todo 解密
//	key := utils.ReadResource("/keys/auth_private-key.der")
//}

type SessionKeyAuthenticator struct {
}

func (SessionKeyAuthenticator) Authenticate(request *AuthenticationRequest) any {
	resp := object.NewComboTokenResJson()
	var loginSuccess bool
	if request.SessionKeyData == nil || request.SessionKeyRequest == nil {
		log.Println("requestData should never be nil")
		return nil
	}
	requestData := request.SessionKeyRequest
	_ = requestData
	loginData := request.SessionKeyData
	address := utils.Address(request.Context)
	// Get account from db.
	account := game.GetAccountById(loginData.UID)
	if account == nil || account.SessionKey != loginData.Token {
		log.Println("user not exit, req come from ", address)
	} else {
		loginSuccess = true
	}
	// todo account.SessionKey != loginData.Token
	loginSuccess = true
	if loginSuccess {
		resp.Message = "OK"
		resp.Data.ComboID = "157795300"
		resp.Data.ComboToken = account.GenerateLoginToken()
	} else {
		resp.Retcode = -200
		resp.Message = "Go-G错误"
	}
	return resp
}
