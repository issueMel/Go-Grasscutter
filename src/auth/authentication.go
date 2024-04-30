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
	var req *object.LoginAccountRequestJson
	req = request.PasswordRequest
	if req == nil {
		log.Fatal("requestData should never be nil")
		return nil
	}
	address := utils.Address(request.Context)
	//responseMessage := translate("messages.dispatch.account.username_error")
	//loggerMessage := ""

	// Get account from db.
	account := game.GetAccountName(req.Account)
	// todo autoCreate
	if account == nil {
		log.Fatal("user not exit, req come from ", address)
		return nil
	}
	obj := &object.LoginResultJson{
		Message: "OK",
		VerifyData: object.VerifyData{
			Account: object.VerifyAccountData{
				UID:   account.ID,
				Email: account.GetEmail(),
				Token: account.GenerateSessionKey(),
			},
		},
	}
	return obj
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
