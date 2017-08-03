package upload

import (
	"gopkg.in/kataras/iris.v6"
	"os"
	"io"
)

func BindRoute(r *iris.Router)  {
	r.Post("/files", iris.LimitRequestBodySize(10<<20),
		func(ctx *iris.Context) {

			file, info, err := ctx.FormFile("image")


			if err != nil {
				ctx.HTML(iris.StatusInternalServerError,
					"Error while uploading: <b>"+err.Error()+"</b>")
				return
			}

			defer file.Close()
			fname := info.Filename

			// Create a file with the same name
			// assuming that you have a folder named 'uploads'
			out, err := os.OpenFile("./public/uploads/"+fname,
				os.O_WRONLY|os.O_CREATE, 0666)

			if err != nil {
				ctx.JSON(iris.StatusInternalServerError,iris.Map{
					"error" : err.Error(),
				})
				return
			}else{
				ctx.JSON(iris.StatusOK,iris.Map{"info":info})
			}
			defer out.Close()

			io.Copy(out, file)
	})
}
