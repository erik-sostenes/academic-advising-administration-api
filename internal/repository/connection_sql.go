package repository

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	onceMySQL sync.Once
	sqlConnection *sql.DB
)

// LoadMySQLConnection create the connection to MYSQL 
func LoadMySQLConnection() (*sql.DB, error) {
	var err error

	onceMySQL.Do(func() {
		driverName := os.Getenv("MYSQL_DRIVER")

		url := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			"root",
			os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("MYSQL_HOST"),
			os.Getenv("MYSQL_PORT"),
			"advisories",
		)
		sqlConnection, err = sql.Open(driverName, url)
		if err != nil {
			return
		}
		err = sqlConnection.Ping()
	})
	return sqlConnection, err
}

// NewMySQL create the MySQL instance 
func NewMySQL() *sql.DB {
	db, err := LoadMySQLConnection()
	if err != nil {
		panic(err)
	}
	return db
}
