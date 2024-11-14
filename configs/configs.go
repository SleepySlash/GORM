package configs

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {
	DB_URI := os.Getenv("DATABASE_URI")

	db, err := gorm.Open(postgres.Open(DB_URI), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
