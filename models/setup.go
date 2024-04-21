package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DataInstance struct {
	Db *gorm.DB
}

	var DB DataInstance

	func ConnectDB() {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error while loading .env files: %v", err)
		}

		host     := os.Getenv("DB_HOST")
		user     := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname   := os.Getenv("DB_NAME")
		port     := os.Getenv("DB_PORT")

		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbname, port,
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		if err != nil {
			log.Fatal("Failed to connect database, \n", err)
			os.Exit(2)
		}

		log.Println("Connected to database")
		db.Logger = logger.Default.LogMode(logger.Info)

		log.Println("Running the programs")
		db.AutoMigrate(&Loketsatu{})

		DB = DataInstance{
			Db: db,
		}
	}