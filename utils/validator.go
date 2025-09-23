package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Validator = validator.New()

func ValidateStruct(s interface{}) []string {
	var errors []string

	err := Validator.Struct(s)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("field %s failed on %s tag", err.Field(), err.Tag())
			errors = append(errors, message)
		}
	}
	return errors
}

func HasUperAndLowerCase(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	return HasUpperCase(password) && HasLowerCase(password)
}
