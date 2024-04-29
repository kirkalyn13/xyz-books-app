package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

// LoadDatabase loads SQLite
func LoadDatabase() error {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		return err
	}

	// CreateTables(db, Tables)
	log.Println("Successfully Loaded Database.")

	DB = db

	return nil
}

// CreateTables translates models into database tables
func CreateTables(db *gorm.DB, tables []interface{}) {
	err := db.AutoMigrate(tables...)

	if err != nil {
		log.Fatalf(fmt.Sprintf("DropTables - Migrator :%s", err.Error()))
	}

	log.Println("Successfully Migrated Tables.")
}
