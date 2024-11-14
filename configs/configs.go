package configs

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
}

func DatabaseConnection() *gorm.DB {
    DB_URI := os.Getenv("DATABASE_URI")
    if DB_URI == "" {
        panic("DATABASE_URI environment variable is not set")
    }

    db, err := gorm.Open(postgres.Open(DB_URI), &gorm.Config{})
    if err != nil {
        panic(fmt.Sprintf("failed to connect to database: %v", err))
    }
    return db
}