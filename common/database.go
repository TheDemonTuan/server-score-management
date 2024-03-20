package common

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DBConn *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_SSL_MODE"), os.Getenv("DB_TIME_ZONE"))
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		//Logger:      logger.Default.LogMode(logger.Silent),
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	//

	if err != nil {
		panic("Database connection failed")
	}
	//log.Println("Connection successfully")

	DBConn = dbConn
	defer runMigrate()
}

func runMigrate() {
	// Drop table
	//if err := DBConn.Migrator().DropTable(&entity.Department{}, &entity.Instructor{}, &entity.Subject{}, &entity.Student{}, &entity.Grade{}, &entity.Class{}, &entity.Assignment{}, &entity.User{}); err != nil {
	//	panic(err)
	//}
	//if err := DBConn.AutoMigrate(&entity.Department{}, &entity.Instructor{}, &entity.Subject{}, &entity.Student{}, &entity.Grade{}, &entity.Class{}, &entity.Assignment{}, &entity.User{}); err != nil {
	//	panic(err)
	//}
	//log.Println("Success to migrate")
}
