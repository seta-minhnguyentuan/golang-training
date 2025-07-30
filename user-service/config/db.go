package config

import (
	"fmt"
	"log"
	"user-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	host := Getenv("DB_HOST", "localhost")
	user := Getenv("DB_USER", "postgres")
	password := Getenv("DB_PASSWORD", "postgres")
	dbname := Getenv("DB_NAME", "user_service")
	port := Getenv("DB_PORT", "5432")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
	db.AutoMigrate(&models.User{})

	return db
}
