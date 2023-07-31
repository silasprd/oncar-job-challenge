package db

import (
	"fmt"
	model "oncar-job-challenge/core/model"

	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func OpenConnection(envPath string) (*gorm.DB, error) {

	err := godotenv.Load(envPath)
	if err != nil {
		return nil, fmt.Errorf("Erro ao carregar variáveis de ambiente: %w", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil

}

func AutoMigrateTables(db *gorm.DB) error {

	err := db.AutoMigrate(&model.Car{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&model.Contact{})
	if err != nil {
		return err
	}

	return nil

}
