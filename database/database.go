package database

import (
	"fmt"
	"github.com/BerkatPS/internal/domain/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get database configuration from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Open a database connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUsername, dbPassword, dbHost, dbPort, dbName)

	// Konfigurasi untuk log mode
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Lambatnya waktu threshold
			LogLevel:      logger.Silent, // Silent, atau Silent
			Colorful:      true,          // Memuat dengan warna
		},
	)

	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(
		&model.Users{},
		&model.Accounts{},
		&model.Categorys{},
		&model.Transaction{},
		&model.MonthlyReport{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err.Error())
	}

	DB = db

	return DB, nil
}
