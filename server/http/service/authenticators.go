package service

import (
	"Go-Grasscutter/game/player"
	"Go-Grasscutter/log"
	"Go-Grasscutter/server/http/object"
	"Go-Grasscutter/utils"
	"Go-Grasscutter/utils/lang"
	"github.com/cloudwego/hertz/pkg/app"
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

// Authenticate Handles the authentication request from the username and password form.
func (PasswordAuthenticator) Authenticate(request *AuthenticationRequest) any {
	resp := object.NewLoginResultJson()

	requestData := request.PasswordRequest
	if requestData == nil {
		log.SugaredLogger.Error("requestData should never be nil")
		return nil
	}
	var successfulLogin bool
	address := utils.Address(request.Context)
	responseMessage := lang.Translate("messages.dispatch.account.username_error")
	// loggerMessage := ""

	// Get account from db.
	account := player.GetAccountName(requestData.Account)
	// todo INCOMPLETE: autoCreate
	successfulLogin = account != nil
	if !successfulLogin {
		log.SugaredLogger.Info(lang.Translate("messages.dispatch.account.account_login_create_error") + address)
	}

	// Set response data.
	if successfulLogin {
		resp.Message = "OK"
		resp.Data.Account.UID = account.ID
		resp.Data.Account.Token = account.GenerateSessionKey()
		resp.Data.Account.Email = account.GetEmail()
	} else {
		resp.RetCode = -201
		resp.Message = responseMessage
	}
	return resp
}

type ExperimentalPasswordAuthenticator struct {
}

func (SessionKeyAuthenticator) ExperimentalPasswordAuthenticator(request *AuthenticationRequest) any {
	return nil
}

type TokenAuthenticator struct {
}

// Authenticate Handles the authentication request from the game when using a registry token
func (TokenAuthenticator) Authenticate(request *AuthenticationRequest) any {
	resp := object.NewLoginResultJson()
	requestData := request.TokenRequest
	if requestData == nil {
		log.SugaredLogger.Error("requestData should never be nil")
		return nil
	}
	var successfulLogin bool
	// todo someone go in system
	//address := utils.Address(request.Context)

	// Get account from database.
	account := player.GetAccountById(requestData.UID)
	// Check if account exists/token is valid.
	successfulLogin = account != nil && account.SessionKey == requestData.Token

	// Set response data.
	if successfulLogin {
		resp.Message = "OK"
		resp.Data.Account.UID = account.ID
		resp.Data.Account.Token = account.GenerateSessionKey()
		resp.Data.Account.Email = account.GetEmail()
	} else {
		resp.RetCode = -201
		resp.Message = lang.Translate("messages.dispatch.account.account_cache_error")

		// Log the failure.
	}
	return resp
}

type SessionKeyAuthenticator struct{}

func (SessionKeyAuthenticator) Authenticate(request *AuthenticationRequest) any {
	resp := object.NewComboTokenResJson()
	var loginSuccess bool
	if request.SessionKeyData == nil || request.SessionKeyRequest == nil {
		log.SugaredLogger.Error("requestData should never be nil")
		return nil
	}
	requestData := request.SessionKeyRequest
	_ = requestData
	loginData := request.SessionKeyData
	address := utils.Address(request.Context)
	// Get account from db.
	account := player.GetAccountById(loginData.UID)

	loginSuccess = account != nil && account.SessionKey == loginData.Token

	if loginSuccess {
		resp.Message = "OK"
		resp.Data.OpenID = account.ID
		resp.Data.ComboID = "157795300"
		resp.Data.ComboToken = account.GenerateLoginToken()
	} else {
		resp.Retcode = -200
		resp.Message = lang.Translate("messages.dispatch.account.session_key_error")
		log.SugaredLogger.Info(address)
	}
	return resp
}

type SessionTokenValidator struct {
}

// todo INCOMPLETE: SessionTokenValidator()
func (SessionTokenValidator) Authenticate(request *AuthenticationRequest) any {
	return nil
}
