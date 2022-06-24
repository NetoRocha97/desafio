package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDb() {
	str := "host=localhost port=5432 user=postgres dbname=lojas sslmode=disable password=postgres"
	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil {
		log.Fatal("error", err)
	}

	db = database
	config, _ := db.DB()
	config.SetConnMaxIdleTime(10)
	config.SetMaxIdleConns(100)
	config.SetConnMaxLifetime(time.Hour)
	RunMigrations(db)
}

func GetDatabase() *gorm.DB{
	return db
}