package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	customLogger "password-manager-backend/logger"
)

var Mysql *gorm.DB

func ConnectMysql() {
	defer func() {
		if err := recover(); err != nil {
			customLogger.LogError("Couldn't connect to MySQL database")
		} else {
			customLogger.LogInfo("Connected to MySQL database")
		}
	}()

	conn, err := gorm.Open(mysql.Open("user:password@tcp(localhost:3306)/db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	Mysql = conn

	err = Mysql.AutoMigrate()
	if err != nil {
		panic(err)
	}
}
