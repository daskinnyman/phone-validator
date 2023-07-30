package serivces

import (
	"phone_validator/pkg/models"
	"phone_validator/pkg/repositories"
	"phone_validator/pkg/serivces/dtos"

	"github.com/nyaruka/phonenumbers"
)

// Step 1: define the struct for the service
type PhoneValidatorService interface {
	ValidatePhoneNumber(req dtos.ValidatePhoneNumberRequest) (bool, error)
}

type phoneValidatorService struct {
	storage repositories.PhoneNumberRepository
}

// Step 2: create the constructor for the service
func CreatePhoneValidatorService(repo repositories.PhoneNumberRepository) PhoneValidatorService {
	return &phoneValidatorService{
		storage: repo,
	}
}

// Step 3: implement the methods for the service
func (p *phoneValidatorService) ValidatePhoneNumber(req dtos.ValidatePhoneNumberRequest) (bool, error) {
	phoneNumberToValidate := req.PhoneNumber

	// 1. Check if phone number is empty or spaces
	if len(phoneNumberToValidate) == 0 || phoneNumberToValidate == " " {
		return false, nil
	}

	// 2. Check if phone number is existing in the database
	phoneRecord, err := p.storage.GetPhoneNumber(phoneNumberToValidate)

	if err != nil {
		return false, err
	}

	// 3. Return true if phone number is existing in the database
	if phoneRecord.PhoneNumber != "" {
		return phoneRecord.Valid, nil
	}

	// 4. parse the phone number
	phoneNumber, err := phonenumbers.Parse(phoneNumberToValidate, "TW")

	if err != nil {
		return false, err
	}

	isValid := phonenumbers.IsValidNumberForRegion(phoneNumber, "TW")

	// 5. Save the phone number to the database
	phoneNumberToSave := models.PhoneNumber{
		PhoneNumber: phoneNumberToValidate,
		Valid:       isValid,
	}

	res, err := p.storage.CreatePhoneNumber(phoneNumberToSave)

	if err != nil {
		return false, err
	}

	return res.Valid, nil
}
