package lib

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/pressly/goose"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db

	return db
}

func RunMigrations(db *sql.DB) {
	log.Println("running migration...")
	err := goose.Up(db, "migrations")
	if err != nil {
		log.Println("Failed run migration : ", err)
	}
	log.Println("migration finish")
}

func InitDatabase() {
	db := connectDB()
	sqlDB, _ := db.DB()
	RunMigrations(sqlDB)
}
