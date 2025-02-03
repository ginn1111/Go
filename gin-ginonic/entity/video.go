package entity

import "time"

type Person struct {
	ID        uint64 `json:"id" gorm:"primary_key;auto_increment"`
	FirstName string `json:"firstname" binding:"required" gorm:"type:varchar(32)"`
	LastName  string `json:"lastname" binding:"required" gorm:"type:varchar(32)"`
	Age       uint8  `json:"age" binding:"required,gte=1,lte=130"`
	Email     string `json:"email" binding:"required,email" gorm:"type:varchar(256)"`
}

type Video struct {
	ID          uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Title       string    `json:"title" binding:"min=2,max=10" gorm:"type:varchar(10)"`
	Url         string    `json:"url" binding:"required,url" gorm:"type:varchar(256);UNIQUE"`
	Description string    `json:"description" binding:"max=20" gorm:"type:varchar(20)"`
	Author      Person    `json:"author" binding:"required" gorm:"foreignkey:PersonID"`
	PersonID    uint64    `json:"-"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (Video) TableName() string {
	return "videos"
}
