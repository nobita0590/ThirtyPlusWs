package auth

import (
	"gopkg.in/kataras/iris.v6"
	"github.com/nobita0590/ThirtyPlusWs/helper"
)

var (
	response helper.ResponseHelper
)

func BindRoute(route *iris.Router)  {
	route.Post("/login",loginAction)
}

func loginAction(ctx *iris.Context)  {
	/*userName := ctx.FormValue("user_name")
	password := ctx.FormValue("password")

	urlData := url.Values{}
	urlData.Set()
	resp, err := http.Post(config.CRMUrlLogin + "/call/index/memberlogin", "application/x-www-form-urlencoded", strings.NewReader(urlData.Encode()))*/
}