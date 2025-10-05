package validation

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidationErrorMessages(err error) map[string]string {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		return generateErrorMessages(validationErrors)
	}
	return nil
}

func generateErrorMessages(validationErrors validator.ValidationErrors) map[string]string {
	errorsMap := make(map[string]string)
	for _, err := range validationErrors {
		fieldName := strings.ToLower(err.Field())
		tag := strings.ToLower(err.Tag())

		customMessage := customMessages[tag]
		if customMessage != "" {
			errorsMap[fieldName] = formatErrorMessage(customMessage, err, tag)
		} else {
			errorsMap[fieldName] = defaultErrorMessage(err)
		}
	}
	return errorsMap
}

func formatErrorMessage(customMessage string, err validator.FieldError, tag string) string {
	errField := strings.ToLower(err.Field())
	if tag == "min" || tag == "max" || tag == "len" {
		return fmt.Sprintf(customMessage, errField, err.Param())
	}
	return fmt.Sprintf(customMessage, errField)
}

func defaultErrorMessage(err validator.FieldError) string {
	return fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", strings.ToLower(err.Field()), err.Tag())
}

func NewValidator() *validator.Validate {
	validator := validator.New()
	if err := validator.RegisterValidation("password", password); err != nil {
		return nil
	}
	return validator
}
