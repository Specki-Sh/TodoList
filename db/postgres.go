package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var database *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "mydb"
)

func createTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		is_done BOOLEAN NOT NULL
	)`
	_, err := db.Exec(query)
	return err
}

func initDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = createTable(db)
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
func GetDBConn() *sql.DB {
	return database
}

func CloseDbConnection() error {
	if err := GetDBConn().Close(); err != nil {
		fmt.Errorf("error occurred on database connection closing: %s", err.Error())
	}
	return nil
}
