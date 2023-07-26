package db

import (
	model "oncar-job-challenge/api/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func OpenConnection() (*gorm.DB, error) {

	dsn := "root:root@tcp(127.0.0.1:3306)/oncar?charset=utf8mb4&parseTime=True&loc=Local"

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

	return nil

}
