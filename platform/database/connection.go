package database

import (
	"database/sql"
	"log"
	"os"
	"sync"

	_ "github.com/joho/godotenv/autoload"
)

var once sync.Once
var connection *sql.DB

func Connection() *sql.DB {
	once.Do(func() {
		connection = initialize()
	})

	return connection
}

func initialize() *sql.DB {
	switch os.Getenv("DB_DRIVER") {
	case "postgresql":
		return initializePostgres()
	case "mysql":
		return initializeMysql()
	default:
		log.Fatalln("You must specify a database driver. Choices are 'postgresql' or 'mysql'")
		return nil
	}
}
