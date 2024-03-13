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
		PrepareStmt: true,
	})

	if err != nil {
		panic("Database connection failed")
	}
	log.Println("Connection successfully")

	DBConn = dbConn
	defer runMigrate()
}

func runMigrate() {
	// Drop table
	//if err := DBConn.Migrator().DropTable(&entity.Department{}, &entity.Teacher{}, &entity.Subject{}, &entity.Student{}, &entity.Transcript{}, &entity.Class{}, &entity.User{}); err != nil {
	//	panic(err)
	//}

	if err := DBConn.AutoMigrate(&entity.Department{}, &entity.Teacher{}, &entity.Subject{}, &entity.Student{}, &entity.Transcript{}, &entity.Class{}, &entity.User{}); err != nil {
		panic(err)
	}
	log.Println("Success to migrate")
}
