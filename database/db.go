package database

import (
	"belajar-middleware/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB
var err error

func StartDB() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Connection Failed to Open %v\n", err)
		panic(err)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Product{})

	fmt.Println("Connection Established")

}

func GetDB() *gorm.DB {
	return db
}
