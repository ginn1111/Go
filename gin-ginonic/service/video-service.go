package service

import (
	"github.com/go/go-goinc/entity"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (vs *videoService) Save(v entity.Video) entity.Video {
	vs.videos = append(vs.videos, v)

	return v
}

func (vs *videoService) FindAll() []entity.Video {
	return vs.videos
}
