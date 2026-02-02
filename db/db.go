package db

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dbName := os.Getenv("DATABASE")
	if dbName == "" {
		dbName = "./test.db"
	}
	return gorm.Open(sqlite.Open(dbName), &gorm.Config{})
}
