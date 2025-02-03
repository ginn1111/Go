package repository

import (
	"github.com/go/go-goinc/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type VideoRepository interface {
	Save(entity.Video)
	Update(entity.Video)
	Delete(videoId uint64) entity.Video
	FindAll() []entity.Video
	Close()
}

type videoRepository struct {
	connection *gorm.DB
}

func NewVideoRepository() VideoRepository {

	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Connection failed!")
	}

	// db.AutoMigrate(&entity.Video{}, &entity.Person{})

	return &videoRepository{
		connection: db,
	}
}

func (vdRepo *videoRepository) Save(video entity.Video) {
	vdRepo.connection.Create(&video)
}

func (vdRepo *videoRepository) Update(video entity.Video) {
	vdRepo.connection.Save(&video)
}

func (vdRepo *videoRepository) Delete(videoId uint64) entity.Video {
	var deletedVideo entity.Video
	vdRepo.connection.Where("ID=?", videoId).Delete(deletedVideo)

	return deletedVideo
}

func (vdRepo *videoRepository) FindAll() []entity.Video {
	var videos []entity.Video

	vdRepo.connection.Debug().Preload("Author").Find(&videos)
	// vdRepo.connection.Debug().Set("gorm:auto_preload", true).Find(&videos)

	return videos
}

func (vdRepo *videoRepository) Close() {
	vdRepo.connection.Close()
}
