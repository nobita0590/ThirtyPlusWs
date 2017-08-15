package news

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/mgo.v2/bson"
	"github.com/nobita0590/ThirtyPlusWs/helper"
)

var (
	response helper.ResponseHelper
)

func BindRoute(route *iris.Router)  {
	/* news api */
	route.Options("/index", func(ctx *iris.Context) {
		ctx.JSON(iris.StatusOK,bson.M{"status":true})
	})
	route.Get("/index", getNews)
	route.Post("/index", insertNews)
	route.Put("/index",updateNews)
	route.Delete("/index",deleteNews)
	route.Get("/index/list",getListNews)

	/* category api */
	route.Options("/category", func(ctx *iris.Context) {
		ctx.JSON(iris.StatusOK,bson.M{"status":true})
	})
	route.Get("/category",getCategory)
	route.Post("/category",insertCategory)
	route.Put("/category",updateCategory)
	route.Delete("/category",deleteCategory)
	route.Get("/category/list",getCategories)
}
