package repositories

import (
	"phone_validator/pkg/models"

	"gorm.io/gorm"
)

// Step 1: define the struct for the repository
type PhoneNumberRepository interface {
	// CreatePhoneNumber creates the phone number
	CreatePhoneNumber(phoneNumber models.PhoneNumber) (models.PhoneNumber, error)
	// GetPhoneNumber gets the phone number
	GetPhoneNumber(phoneNumber string) (models.PhoneNumber, error)
}

type phoneNumberRepository struct {
	db *gorm.DB
}

// Step 2: create the constructor for the repository
func CreatePhoneNumberRepository(db *gorm.DB) PhoneNumberRepository {
	return &phoneNumberRepository{
		db: db,
	}
}

// Step 3: implement the methods for the repository
func (r *phoneNumberRepository) CreatePhoneNumber(phoneNumber models.PhoneNumber) (models.PhoneNumber, error) {
	err := r.db.Create(&phoneNumber).Error

	if err != nil {
		return phoneNumber, err
	}

	return phoneNumber, nil
}

func (r *phoneNumberRepository) GetPhoneNumber(phoneNumber string) (models.PhoneNumber, error) {
	var phoneNumberRes models.PhoneNumber
	r.db.First(&phoneNumberRes, phoneNumber)

	return phoneNumberRes, nil
}
