package dispatch

import (
	"Go-Grasscutter/src/auth"
	"Go-Grasscutter/src/server/http/object"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"net/http"
)

// {"account":"user",
//
//	"password":"Byctjr3n+Gz2DOnr1T8BrruNywK4fxM5iQZjVqbffaEY8N9nVNoKxJiZ1hSGI4Q+1uzsaxI33mQC5A2FW1rLJBGuPxJhHPLDut4xSKhwBr0OCjGTVG8R9P3yrzImUs4p9q3CuBUck9/MKL22R41Ch3/ep297mMHymVLxH6c5y3c=",
//	"is_crypto":true}
func ClientLogin(c context.Context, ctx *app.RequestContext) {
	var req *object.LoginAccountRequestJson
	// Validate data.
	err := ctx.BindJSON(&req)
	if err != nil {
		log.Println(err)
		return
	}

	// Pass data to authentication handler.
	var pas = new(auth.PasswordAuthenticator)
	resp := pas.Authenticate(&auth.AuthenticationRequest{
		Context:         ctx,
		PasswordRequest: req,
	}).(*object.LoginResultJson)
	fmt.Printf("%+v\n", resp.VerifyData)
	ctx.JSON(http.StatusOK, resp)
}
