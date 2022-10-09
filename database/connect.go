package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", Config("DB_USERNAME"), Config("DB_PASSWORD"), Config("DB_HOST"), Config("DB_PORT"), Config("DB_DATABASE"))
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	return db, err
}

func Config(key string) string {
	return os.Getenv(key)
}
