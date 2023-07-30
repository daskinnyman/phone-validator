package controllers

import (
	"log"
	"net/http"
	"phone_validator/pkg/serivces"
	"phone_validator/pkg/serivces/dtos"

	"github.com/gin-gonic/gin"
)

// Step 1: define the struct for the controller
type PhoneValidatorHandler struct {
	PhoneValidatorService serivces.PhoneValidatorService
}

// Step 2: create the constructor for the controller
func CreatePhoneValidatorHandler(e *gin.Engine, phoneValidatorService serivces.PhoneValidatorService) {
	handler := &PhoneValidatorHandler{
		PhoneValidatorService: phoneValidatorService,
	}

	e.POST("/api/validate", handler.ValidatePhoneNumber)
}

// Step 3: implement the methods for the controller
func (handler *PhoneValidatorHandler) ValidatePhoneNumber(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	var phoneToValidate dtos.ValidatePhoneNumberRequest
	// 1. Get the request body
	err := c.ShouldBindJSON(&phoneToValidate)

	if err != nil {
		log.Printf("handler error: %v", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	// 2. Call the service
	isValid, err := handler.PhoneValidatorService.ValidatePhoneNumber(phoneToValidate)

	if err != nil {
		log.Printf("handler error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	// 3. Return the response
	c.JSON(http.StatusOK, gin.H{
		"error": nil,
		"valid": isValid,
	})
}
