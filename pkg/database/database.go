package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDBConnection() {
	var err error

	var DB_HOST = os.Getenv("DB_HOST")
	var DB_USER = os.Getenv("DB_USER")
	var DB_NAME = os.Getenv("DB_NAME")
	var DB_PORT = os.Getenv("DB_PORT")

	// var DB_HOST = "localhost"
	// var DB_USER = "postgres"
	// var DB_NAME = "go_market_local"
	// var DB_PASSWORD = "123"
	// var DB_PORT = "5432"

	// dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s", DB_HOST, DB_USER, DB_NAME, DB_PASSWORD, DB_PORT)
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s", DB_HOST, DB_USER, DB_NAME, DB_PORT)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Database")
}
