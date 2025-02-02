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
	"github.com/go/go-goinc/entity"
	"github.com/go/go-goinc/middlewares"
	"github.com/go/go-goinc/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)

	loginService    service.LoginService       = service.NewLoginService()
	jwtService      service.JWTService         = service.NewJWTService()
	loginController controller.LoginController = controller.NewLoginController(jwtService, loginService)
)

func setupOutputFile() {
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
	setupOutputFile()
	registerValidators()

	r := gin.New()

	r.Use(gin.Recovery(), middlewares.Logger())

	r.Static("/css", "./template/css")
	r.LoadHTMLGlob("template/*.html")

	r.POST("/login", func(ctx *gin.Context) {

		var credentials entity.Credentials

		err := ctx.ShouldBindJSON(&credentials)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		}

		token := loginController.Login(credentials)

		defer func() {
			err := recover()

			ctx.JSON(http.StatusUnauthorized, err)
		}()

		if token == "" {
			ctx.JSON(http.StatusUnauthorized, "")

			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	})

	apiRoutes := r.Group("/api", middlewares.JWTAuth())
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

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
