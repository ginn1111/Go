package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go/go-goinc/entity"
	"github.com/go/go-goinc/service"
)

type VideoController interface {
	Save(c *gin.Context) error
	FindAll() []entity.Video
	ShowAll(c *gin.Context)
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	if err := ctx.ShouldBindJSON(&video); err != nil {
		return err
	}

	c.service.Save(video)

	return nil
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.FindAll()

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"videos": videos,
	})

}
