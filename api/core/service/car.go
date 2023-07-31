package service

import (
	model "oncar-job-challenge/core/model"

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
	err := s.db.Preload("Contacts").Find(&cars).Error
	if err != nil {
		return nil, err
	}
	return cars, err
}

func (s *CarService) GetCarByID(id uint) (*model.Car, error) {
	var car model.Car
	err := s.db.Preload("Contacts").First(&car, id).Error
	if err != nil {
		return nil, err
	}
	return &car, nil
}

func (s *CarService) DeleteCar(id uint) error {
	var car model.Car
	result := s.db.Delete(&car, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
