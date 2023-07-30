package models

// Create gorm model for phone number and it's validation result
type PhoneNumber struct {
	PhoneNumber string `gorm:"primaryKey" json:"phone_number"`
	Valid       bool
	UpdatedAt   int
}
