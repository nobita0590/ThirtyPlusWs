package news

import (
	"gopkg.in/kataras/iris.v6"
	"github.com/nobita0590/ThirtyPlusWs/modules/shared"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

func BindRoute(route *iris.Router)  {
	route.Get("/index", func(ctx *iris.Context) {
		fmt.Println("news/index")
		if db,ok := shared.GetDbSession(ctx);ok{
			err := db.DB("test").C("test").Insert(bson.M{"name":"test"})
			ctx.JSON(iris.StatusOK,bson.M{"error":err})
		}else{
			ctx.JSON(iris.StatusOK,bson.M{"error":"can not get db session"})
		}
	})
}
