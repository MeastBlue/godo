package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetDatabase() (*sqlx.DB, error) {
	driver, url := getDatabaseString()
	return sqlx.Connect(driver, url)
}

func getDatabaseString() (string, string) {
	driver := os.Getenv("db.Driver")
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("db.Host"), os.Getenv("db.Port"), os.Getenv("db.User"),
		os.Getenv("db.Password"), os.Getenv("db.Name"), os.Getenv("db.SSL"))

	return driver, url
}
