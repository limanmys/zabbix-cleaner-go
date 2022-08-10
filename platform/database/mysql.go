package database

import (
	"database/sql"
	"fmt"
	"os"
)

func initializeMysql() *sql.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	connection, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil
	}
	return connection
}
