package auth

import (
	"Go-Grasscutter/src/database"
	"Go-Grasscutter/src/server/http/object"
	"Go-Grasscutter/utils"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
)

func PasswordAuthenticator(c context.Context, ctx *app.RequestContext) {
	var req *object.LoginAccountRequestJson
	err := ctx.Bind(req)
	if err != nil {
		log.Fatal("requestData should never be nil")
		return
	}
	address := utils.Address(ctx)
	//responseMessage := translate("messages.dispatch.account.username_error")
	//loggerMessage := ""

	// Get account from database.
	account := database.GetAccountName(req.Account)
	// todo autoCreate
	if account == nil {
		log.Fatal("user not exit, req come from ", address)
		return
	}
	ctx.JSON(200, &object.LoginResultJson{
		Message: "OK",
		RetCode: 0,
		Data: object.VerifyData{
			Account: object.VerifyAccountData{
				UID:   account.ID,
				Email: account.GetEmail(),
				Token: account.GenerateSessionKey(),
			},
		},
	})
}
