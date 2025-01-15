package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	// c, _ := mysql.NewConnector(&mysql.Config{
	// 	User:   "root",
	// 	Passwd: "root",
	// 	Addr:   "127.0.0.1",
	// 	DBName: "bookstore",
	// })
	//

	d, err := gorm.Open("mysql", "root:@/bookstore?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	db = d

	log.Println("DB connected")
}

func GetDB() *gorm.DB {
	return db
}
