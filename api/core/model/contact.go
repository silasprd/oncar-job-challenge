package models

type Contact struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"not null"`
	Email string `gorm:"not null"`
	Phone string `gorm:"not null"`
	CarID uint
}
