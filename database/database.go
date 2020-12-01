package database

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func GetDatabase() (*sqlx.DB, error) {
	driver, url := getDatabaseString()
	return sqlx.Connect(driver, url)
}

func getDatabaseString() (string, string) {
	driver := os.Getenv("db.Driver")
	url := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true",
		os.Getenv("db.User"), os.Getenv("db.Password"), os.Getenv("db.Host"),
		os.Getenv("db.Port"), os.Getenv("db.Name"))

	return driver, url
}
