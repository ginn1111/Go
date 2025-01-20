package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go/go-goinc/controller"
	"github.com/go/go-goinc/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	r := gin.Default()

	r.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})

	r.POST("/video", func(ctx *gin.Context) {
		videoController.Save(ctx)
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
