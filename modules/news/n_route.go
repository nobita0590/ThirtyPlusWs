package news

import (
	"gopkg.in/kataras/iris.v6"
	"github.com/nobita0590/ThirtyPlusWs/modules/shared"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"github.com/nobita0590/ThirtyPlusWs/models"
	"github.com/nobita0590/ThirtyPlusWs/helper"
)

var (
	response helper.ResponseHelper
)

func BindRoute(route *iris.Router)  {
	// news api
	route.Get("/index", func(ctx *iris.Context) {
		fmt.Println("news/index")
		if db,ok := shared.GetDbSession(ctx);ok{
			err := db.DB("test").C("test").Insert(bson.M{"name":"test"})
			ctx.JSON(iris.StatusOK,bson.M{"error":err.Error()})
		}else{
			ctx.JSON(iris.StatusOK,bson.M{"error":"can not get db session"})
		}
	})
	route.Post("/index", func(ctx *iris.Context) {
		if db,ok := shared.GetDbSession(ctx);ok{
			newsModel := models.NewMainModel(db).GetNewsModel()
			news := models.News{}
			if err := ctx.ReadForm(&news);err == nil {
				if err := newsModel.Insert(&news);err == nil {
					response.Success(ctx,bson.M{"data":news.Id})
				} else {
					ctx.JSON(iris.StatusBadRequest,bson.M{"error":err.Error()})
				}
			}else {
				ctx.JSON(iris.StatusBadRequest,bson.M{
					"error":err.Error(),
					"form": ctx.Request.Form,
				})
			}
		}else {
			ctx.JSON(iris.StatusBadRequest,bson.M{"error":"error connect db"})
		}
	})
	route.Options("/index", func(ctx *iris.Context) {
		ctx.JSON(iris.StatusOK,bson.M{"status":true})
	})
	route.Options("/category", func(ctx *iris.Context) {
		ctx.JSON(iris.StatusOK,bson.M{"status":true})
	})
	// category api
	route.Get("/category",getCategory)
	route.Post("/category",insertCategory)
	route.Put("/category",updateCategory)
	route.Delete("/category",deleteCategory)
	route.Get("/category/list",getCategories)
}


