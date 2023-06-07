package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "mydb"
)

func initDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

// StartDbConnection Creates connection to database
func StartDbConnection() {
	database = initDB()
}

// GetDBConn func for getting db conn globally
func GetDBConn() *gorm.DB {
	return database
}

func CloseDbConnection() error {
	db, err := database.DB()
	if err != nil {
		return fmt.Errorf("error occurred on database connection closing: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		return fmt.Errorf("error occurred on database connection closing: %s", err.Error())
	}
	return nil
}

