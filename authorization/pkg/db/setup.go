package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var Database *gorm.DB

func InitDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/ApolloFederation?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err!=nil{
		log.Println("Connection Failed to Open")
	}
	log.Println("Connection Established")
	Database = db
	return db
}


