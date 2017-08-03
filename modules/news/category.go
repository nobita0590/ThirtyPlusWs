package news

import (
	"gopkg.in/kataras/iris.v6"
	"github.com/nobita0590/ThirtyPlusWs/modules/shared"
	"github.com/nobita0590/ThirtyPlusWs/models"
	"gopkg.in/mgo.v2/bson"
	"github.com/nobita0590/ThirtyPlusWs/my_error"
)

func insertCategory(ctx *iris.Context)  {
	if db,ok := shared.GetDbSession(ctx);ok{
		categoryModel := models.NewMainModel(db).GetCategoryModel()
		category := models.Category{}
		if err := ctx.ReadForm(&category);err == nil {
			if err := categoryModel.Insert(&category);err == nil {
				response.Success(ctx,category.Id)
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

func updateCategory(ctx *iris.Context)  {
	if db,ok := shared.GetDbSession(ctx);ok{
		categoryModel := models.NewMainModel(db).GetCategoryModel()
		//category := models.Category{}
		form := struct {
			models.Category
			Fields []string
		}{}
		if err := ctx.ReadForm(&form);err == nil {
			if len(form.Fields) == 0 {
				err = categoryModel.Update(form.Category)
			}else{
				err = categoryModel.UpdatePartial(form.Category,form.Fields...)
			}
			if err == nil {
				response.Success(ctx,form.Category.Id)
			} else {
				response.ErrorInternalServer(ctx,err)
			}
		}else {
			ctx.JSON(iris.StatusBadRequest,bson.M{
				"error":err.Error(),
				"form": ctx.Request.Form,
			})
		}
	} else {
		response.ErrorInternalServer(ctx,my_error.DatabaseError)
	}
}


func getCategory(ctx *iris.Context) {
	filter := models.CategoryFilter{}
	if err := ctx.ReadForm(&filter);err == nil {
		if db,ok := shared.GetDbSession(ctx);ok {
			categoryModel := models.NewMainModel(db).GetCategoryModel()
			if category,err := categoryModel.Get(filter);err == nil {
				response.Success(ctx,category)
			}else{
				response.ErrorInternalServer(ctx,err)
			}
		}else{
			response.ErrorInternalServer(ctx,my_error.DatabaseError)
		}
	}else{
		response.ErrorBadRequest(ctx,err)
	}
}

func getCategories(ctx *iris.Context)  {
	filter := models.CategoryFilter{}
	if err := ctx.ReadForm(&filter);err == nil {
		if db,ok := shared.GetDbSession(ctx);ok {
			categoryModel := models.NewMainModel(db).GetCategoryModel()
			if categories,err := categoryModel.GetList(filter);err == nil {
				response.Success(ctx,categories)
			}else{
				response.ErrorInternalServer(ctx,err)
			}
		}else{
			response.ErrorInternalServer(ctx,my_error.DatabaseError)
		}
	}else{
		response.ErrorBadRequest(ctx,err)
	}
}



