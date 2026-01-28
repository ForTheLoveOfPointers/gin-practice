package db

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dbName := os.Getenv("DATABASE")
	if dbName == "" {
		dbName = "dev-db-gingonic"
	}

	return gorm.Open(sqlite.Open(dbName), &gorm.Config{})
}
