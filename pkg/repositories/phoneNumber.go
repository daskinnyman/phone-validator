package repositories

import (
	models "phone_validator/pkg/db"

	"phone_validator/pkg/repositories/entities"

	"gorm.io/gorm"
)

// Step 1: define the struct for the repository
type PhoneNumberRepository interface {
	// create a validation result with the phoneNumber
	CreatePhoneNumberResult(phoneNumber entities.PhoneNumber) (entities.PhoneNumber, error)
	// get a validation result of the phoneNumber
	GetPhoneNumberResult(phoneNumber string) (entities.PhoneNumber, error)
}

type phoneNumberRepository struct {
	db *gorm.DB
}

// // Step 2: create the constructor for the repository
func NewPhoneNumberResultRepository(db *gorm.DB) PhoneNumberRepository {
	return &phoneNumberRepository{
		db: db,
	}
}

// Step 3: implement the methods for the repository
func (r *phoneNumberRepository) CreatePhoneNumberResult(phoneNumber entities.PhoneNumber) (entities.PhoneNumber, error) {
	phoneNumberToCreate := models.PhoneNumber{
		PhoneNumber: phoneNumber.PhoneNumber,
		Valid:       phoneNumber.Valid,
		UpdatedAt:   phoneNumber.UpdatedAt,
	}

	err := r.db.Create(&phoneNumberToCreate).Error

	if err != nil {
		return phoneNumber, err
	}

	var phoneNumberEntity entities.PhoneNumber
	phoneNumberEntity.PhoneNumber = phoneNumberToCreate.PhoneNumber
	phoneNumberEntity.Valid = phoneNumberToCreate.Valid
	phoneNumberEntity.UpdatedAt = phoneNumberToCreate.UpdatedAt

	return phoneNumberEntity, nil
}

func (r *phoneNumberRepository) GetPhoneNumberResult(phoneNumber string) (entities.PhoneNumber, error) {
	var phoneNumberRes models.PhoneNumber
	r.db.First(&phoneNumberRes, phoneNumber)

	var phoneNumberEntity entities.PhoneNumber
	phoneNumberEntity.PhoneNumber = phoneNumberRes.PhoneNumber
	phoneNumberEntity.Valid = phoneNumberRes.Valid
	phoneNumberEntity.UpdatedAt = phoneNumberRes.UpdatedAt

	return phoneNumberEntity, nil
}
