package db

import (
	model "oncar-job-challenge/api/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() (*gorm.DB, error) {

	dsn := "root:TelLink1020@tcp(127.0.0.1:5171)/oncar?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.Car{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
