package models

type Car struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Brand    string `gorm:"not null"`
	Model    string `gorm:"not null"`
	Year     int    `gorm:"not null;gt=0"`
	Price    float64
	Contacts []Contact `gorm:"foreignKey:CarID;"`
}
