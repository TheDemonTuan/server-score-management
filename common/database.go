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
	err := DBConn.AutoMigrate(&entity.Department{}, &entity.Teacher{}, &entity.Subject{}, &entity.Student{}, &entity.Transcript{}, &entity.Class{}, &entity.User{})
	if err != nil {
		panic(err)
	}
	log.Println("Success to migrate")
}
