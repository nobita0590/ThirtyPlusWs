package helper

import (
	"gopkg.in/kataras/iris.v6"
)

type ResponseHelper struct {

}
/* group ok */
func (rp ResponseHelper) Success(ctx *iris.Context,data interface{})  {
	ctx.JSON(iris.StatusOK,iris.Map{
		"status" : "Ok",
		"data" : data,
	})
}

func (rp ResponseHelper) DataInvalid(ctx *iris.Context,data interface{})  {
	ctx.JSON(iris.StatusOK,iris.Map{
		"status" : "Invalid",
		"message" : data,
	})
}

/* group error */
func (rp ResponseHelper) ErrorBadRequest(ctx *iris.Context,err error) {
	ctx.JSON(iris.StatusBadRequest,iris.Map{
		"status":"Error",
		"message":err.Error(),
	})
}

func (rp ResponseHelper) ErrorInternalServer(ctx *iris.Context,err error) {
	ctx.JSON(iris.StatusInternalServerError,iris.Map{
		"status":"Error",
		"message":err.Error(),
	})
}

func (rp ResponseHelper) ErrorUnauthorized(ctx *iris.Context,err error) {
	ctx.JSON(iris.StatusUnauthorized,iris.Map{
		"status":"Error",
		"message":err.Error(),
	})
}