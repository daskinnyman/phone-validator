package repositories

import "gorm.io/gorm"

// Step 1: define the struct for the repository
type PhoneNumberRepository interface {
	// ValidatePhoneNumber validates the phone number
	ValidatePhoneNumber(phoneNumber string) (bool, error)
	// GetPhoneNumber gets the phone number
	GetPhoneNumber(phoneNumber string) (bool, error)
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
func (r *phoneNumberRepository) ValidatePhoneNumber(phoneNumber string) (bool, error) {
	panic("implement me")
}

func (r *phoneNumberRepository) GetPhoneNumber(phoneNumber string) (bool, error) {
	panic("implement me")
}
