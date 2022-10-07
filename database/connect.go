package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {

	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", Config("DB_USERNAME"), Config("DB_PASSWORD"), Config("DB_HOST"), Config("DB_PORT"), Config("DB_DATABASE"))
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	return db, err
}

func Config(key string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No Config detected")
	}
	return os.Getenv(key)
}
