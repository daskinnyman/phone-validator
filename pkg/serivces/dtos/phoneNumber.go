package dtos

// Define the request/response for the API
type ValidatePhoneNumberRequest struct {
	PhoneNumber string `json:"phone_number"`
}
