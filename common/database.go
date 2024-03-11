package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"qldiemsv/models/entity"
)

var DBConn *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DB")
	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic("Database connection failed")
	}
	log.Println("Connection successfully")

	DBConn = dbConn
	defer runMigrate()
}

func runMigrate() {
	err := DBConn.AutoMigrate(&entity.Teacher{}, &entity.Subject{}, &entity.User{}, &entity.Department{}, &entity.Class{}, &entity.Transcript{}, &entity.Student{})
	if err != nil {
		panic(err)
	}
	log.Println("Success to migrate")
}
