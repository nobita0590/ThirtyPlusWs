package news

import (
	"gopkg.in/kataras/iris.v6"
	"github.com/nobita0590/ThirtyPlusWs/modules/shared"
	"github.com/nobita0590/ThirtyPlusWs/models"
	"gopkg.in/mgo.v2/bson"
	"github.com/nobita0590/ThirtyPlusWs/my_error"
)

func getNews(ctx *iris.Context) {
	filter := models.NewsFilter{}
	if err := ctx.ReadForm(&filter);err == nil {
		if db,ok := shared.GetDbSession(ctx);ok {
			newsModel := models.NewMainModel(db).GetNewsModel()
			if news,err := newsModel.Get(filter);err == nil {
				response.Success(ctx,news)
			} else {
				response.ErrorInternalServer(ctx,err)
			}
		}else{
			response.ErrorInternalServer(ctx,my_error.DatabaseError)
		}
	}else{
		response.ErrorBadRequest(ctx,err)
	}
}

func insertNews(ctx *iris.Context) {
	if db,ok := shared.GetDbSession(ctx); ok {
		newsModel := models.NewMainModel(db).GetNewsModel()
		news := models.News{}
		if err := ctx.ReadForm(&news);err == nil {
			if err := newsModel.Insert(&news);err == nil {
				response.Success(ctx,news.Id)
			} else {
				response.ErrorInternalServer(ctx,err)
			}
		} else {
			response.ErrorBadRequest(ctx,err)
		}
	} else {
		response.ErrorInternalServer(ctx,my_error.DatabaseError)
	}
}

func updateNews(ctx *iris.Context)  {
	if db,ok := shared.GetDbSession(ctx);ok{
		newsModel := models.NewMainModel(db).GetNewsModel()
		//category := models.Category{}
		form := struct {
			models.News
			Fields      models.Fields
		}{}
		if err := ctx.ReadForm(&form);err == nil {
			if len(form.Fields) == 0 {
				err = newsModel.Update(form.News)
			}else{
				form.Fields.RemoveField("IsActive")
				err = newsModel.UpdatePartial(form.News,form.Fields...)
			}
			if err == nil {
				response.Success(ctx,form.News.Id)
			} else {
				response.ErrorInternalServer(ctx,err)
			}
		}else {
			response.ErrorBadRequest(ctx,err)
		}
	} else {
		response.ErrorInternalServer(ctx,my_error.DatabaseError)
	}
}

func deleteNews(ctx *iris.Context)  {
	if db,ok := shared.GetDbSession(ctx);ok {
		filter := models.NewsFilter{}
		if err := ctx.ReadForm(&filter);err == nil {
			newsModel := models.NewMainModel(db).GetNewsModel()
			if err := newsModel.Delete(filter);err == nil {
				response.Success(ctx,filter.Ids)
			}else{
				response.ErrorInternalServer(ctx,err)
			}
		}else {
			response.ErrorBadRequest(ctx,err)
		}
	} else {
		response.ErrorInternalServer(ctx,my_error.DatabaseError)
	}
}

func getListNews(ctx *iris.Context)  {
	filter := models.NewsFilter{}
	if err := ctx.ReadForm(&filter);err == nil {
		if db,ok := shared.GetDbSession(ctx);ok {
			newsModel := models.NewMainModel(db).GetNewsModel()
			if news,count,err := newsModel.GetList(filter);err == nil {
				response.Success(ctx,bson.M{
					"news" : news,
					"count" : count,
				})
			}else{
				response.ErrorInternalServer(ctx,err)
			}
		}else{
			response.ErrorInternalServer(ctx,my_error.DatabaseError)
		}
	}else {
		response.ErrorBadRequest(ctx, err)
	}
}
