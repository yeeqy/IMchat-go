package config

import (
	"IM-chat/model"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	logging "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

func Init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		logging.Info(err)
	}
	MysqlPath := LoadMysql(file)
	model.ConnDB(MysqlPath)
}

func LoadMysql(file *ini.File) string {
	//DB := file.Section("mysql").Key("DB").String()
	DBHost := file.Section("mysql").Key("DBHost").String()
	DBPort := file.Section("mysql").Key("DBPort").String()
	DBUser := file.Section("mysql").Key("DBUser").String()
	DBPassword := file.Section("mysql").Key("DBPassword").String()
	DBName := file.Section("mysql").Key("DBName").String()
	path := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBPort, DBName)
	return path
}
