package modules

import (
	"gopkg.in/kataras/iris.v6"
	"github.com/nobita0590/ThirtyPlusWs/modules/news"
	"github.com/nobita0590/ThirtyPlusWs/modules/shared"
	"github.com/nobita0590/ThirtyPlusWs/modules/upload"
	"github.com/nobita0590/ThirtyPlusWs/modules/auth"
)

func BindRoute(app *iris.Framework)  {
	app.UseFunc(shared.ShareAllMidlleware)
	newsRoute := app.Party("/news")
	news.BindRoute(newsRoute)

	uploadRoute := app.Party("/upload")
	upload.BindRoute(uploadRoute)

	authRoute := app.Party("/auth")
	auth.BindRoute(authRoute)
}
