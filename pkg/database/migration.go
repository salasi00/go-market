package database

import (
	"fmt"

	"github.com/salasi00/go-market/internal/models"
)

func RunMigration() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Transaction{},
	)

	if err != nil {
		fmt.Println(err)
		panic("migration failed")
	}

	fmt.Println("Migration Success")
}
