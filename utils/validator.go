package utils

import (
	"errors"
	"unicode"

	"github.com/go-playground/validator/v10"
)

// Validator instance
var validate = validator.New()

// Custom validator untuk password tanpa lookahead regex
func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	var (
		hasMinLen  = len(password) >= 8
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	// Cek setiap karakter dalam password
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

// Register validator custom
func InitValidator() {
	_ = validate.RegisterValidation("strong_password", ValidatePassword)
}

// ValidateStruct digunakan untuk memvalidasi struct sebelum diproses
func ValidateStruct(s interface{}) error {
	err := validate.Struct(s)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorMessages []string

		for _, fieldErr := range validationErrors {
			errorMessages = append(errorMessages, fieldErr.Error())
		}

		return errors.New("validation failed: " + errorMessages[0])
	}
	return nil
}
