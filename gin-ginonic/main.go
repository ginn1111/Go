package main

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/go/go-goinc/controller"
	"github.com/go/go-goinc/middlewares"
	"github.com/go/go-goinc/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupOuputFile() {
	f, _ := os.Create("gin.log")

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func registerValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("notoneof", func(fl validator.FieldLevel) bool {
			values := strings.Split(fl.Param(), " ")

			for _, v := range values {
				if v == fl.Field().String() {
					return false
				}
			}

			return true
		})
	}
}

func main() {
	setupOuputFile()
	registerValidators()

	r := gin.New()

	r.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	r.Static("/css", "./template/css")
	r.LoadHTMLGlob("template/*.html")

	apiRoutes := r.Group("/api")
	{

		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, videoController.FindAll())
		})

		apiRoutes.POST("/video", func(ctx *gin.Context) {
			if err := videoController.Save(ctx); err != nil {
				ctx.JSON(http.StatusBadRequest, err.Error())
			} else {

				ctx.JSON(http.StatusOK, videoController.Save(ctx))
			}
		})
	}

	viewRoutes := r.Group("/view")
	{
		viewRoutes.GET("/videos", func(ctx *gin.Context) {
			videoController.ShowAll(ctx)
		})
	}

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
