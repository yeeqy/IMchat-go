package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	logging "github.com/sirupsen/logrus"
)

var DB *gorm.DB

func ConnDB(dsn string) {
	//dsn := "root:yqy000129@tcp(127.0.0.1:3306)/im_chat?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		logging.Info(err)
	}
	db.LogMode(true)
	db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}, &Info{})
	DB = db
}
