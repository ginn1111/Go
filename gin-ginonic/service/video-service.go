package service

import (
	"github.com/go/go-goinc/entity"
	"github.com/go/go-goinc/repository"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	Update(entity.Video)
	Delete(uint64)
	FindAll() []entity.Video
}

type videoService struct {
	repo repository.VideoRepository
}

func New(repo repository.VideoRepository) VideoService {
	return &videoService{
		repo,
	}
}

func (vs *videoService) Save(v entity.Video) entity.Video {
	vs.repo.Save(v)

	return v
}

func (vs *videoService) FindAll() []entity.Video {
	return vs.repo.FindAll()
}

func (vs *videoService) Update(v entity.Video) {
	vs.repo.Update(v)
}
func (vs *videoService) Delete(videoId uint64) {
	vs.repo.Delete(videoId)
}
