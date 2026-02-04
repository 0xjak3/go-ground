package main

import (
	"log"
	"os"
	"xdb/models"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}

// func main() {
// 	AutoMigrate(InitDB())
// }

func main() {
	m, err := migrate.New(
		"file://migrations",
		"postgres://admin:admin@localhost:5432/xdb?sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}

	// รัน Migration ไปยังเวอร์ชันล่าสุด
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
