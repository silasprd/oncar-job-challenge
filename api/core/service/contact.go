package service

import (
	model "oncar-job-challenge/core/model"

	"gorm.io/gorm"
)

type ContactService struct {
	db         *gorm.DB
	carService *CarService
}

func NewContactService(db *gorm.DB, carService *CarService) *ContactService {
	return &ContactService{
		db:         db,
		carService: carService,
	}
}

func (s *ContactService) AddContact(contact model.Contact) error {

	if contact.CarID != 0 {

		car, err := s.carService.GetCarByID(contact.CarID)
		if err != nil {
			return err
		}

		if err := s.db.Create(&contact).Error; err != nil {
			return err
		}

		// Adiciona o contato a lista de contatos do carro
		car.Contacts = append(car.Contacts, contact)
		if err := s.db.Save(&car).Error; err != nil {
			return err
		}

	}

	return nil

}

func (s *ContactService) GetAllContact() ([]model.Contact, error) {

	var contacts []model.Contact
	err := s.db.Find(&contacts).Error
	if err != nil {
		return nil, err
	}

	return contacts, nil

}

func (s *ContactService) GetContactByID(contactID uint) (model.Contact, error) {

	var contact model.Contact
	err := s.db.First(&contact, contactID).Error
	if err != nil {
		return contact, err
	}

	return contact, nil

}

func (s *ContactService) DeleteContact(contactID uint) error {

	var contact model.Contact
	result := s.db.Delete(&contact, contactID)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil

}
