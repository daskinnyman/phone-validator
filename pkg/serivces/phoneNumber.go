package serivces

import (
	"phone_validator/pkg/repositories"
	"phone_validator/pkg/serivces/dtos"
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
	return false, nil
}
