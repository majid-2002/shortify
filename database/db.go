package database

import (
	"log"
	"shortify/models" 
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

// InitDB initializes the database connection and returns the *gorm.DB object
func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=shortify_majid password=shortify@majid dbname=shortify_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
		return nil, err
	}

	err = db.AutoMigrate(&models.URL{})
	if err != nil {
		log.Fatal("Error automigrating the database: ", err)
		return nil, err
	}

	log.Println("Database connection established and models migrated successfully.")
	return db, nil
}
