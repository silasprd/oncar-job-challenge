package service

import (
	model "oncar-job-challenge/api/model"

	"gorm.io/gorm"
)

type CarService struct {
	db *gorm.DB
}

func NewCarService(db *gorm.DB) *CarService {
	return &CarService{db: db}
}

func (s *CarService) AddCar(car model.Car) error {
	return s.db.Create(&car).Error
}

func (s *CarService) GetAllCars() ([]model.Car, error) {
	var cars []model.Car
	err := s.db.Find(&cars).Error
	return cars, err
}

func (s *CarService) GetCarByID(id string) (*model.Car, error) {
	var car model.Car
	err := s.db.First(&car, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &car, nil
}
