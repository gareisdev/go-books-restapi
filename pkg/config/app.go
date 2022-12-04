package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {

	connection_url := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	connection, err := gorm.Open(mysql.Open(connection_url))

	if err != nil {
		log.Fatal("[ERROR] Connection DB failed")
	} else {
		log.Println("[SUCCESS] Connected to the DB")
	}

	db = connection
}

func GetConnectionDB() *gorm.DB {
	return db
}
