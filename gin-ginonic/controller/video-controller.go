package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go/go-goinc/entity"
	"github.com/go/go-goinc/service"
)

type VideoController interface {
	Save(c *gin.Context) entity.Video
	FindAll() []entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)

	c.service.Save(video)

	return video
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}
