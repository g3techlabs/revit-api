package validation

import (
	"errors"
	"fmt"
	"strings"

	authInput "github.com/g3techlabs/revit-api/src/core/auth/input"
	"github.com/go-playground/validator/v10"
)

type IValidator interface {
	Validate(data any) error
}

type CustomValidator struct {
	validator *validator.Validate
}

func NewValidator() IValidator {
	validator := validator.New()

	validator.RegisterStructValidation(IdentifierTypeValidation, authInput.LoginCredentials{})
	validator.RegisterStructValidation(IdentifierTypeValidation, authInput.Identifier{})
	if err := validator.RegisterValidation("password", Password); err != nil {
		return nil
	}
	if err := validator.RegisterValidation("profilepic", ProfilePic); err != nil {
		return nil
	}

	return &CustomValidator{
		validator: validator,
	}
}

func (v CustomValidator) Validate(data any) error {
	return v.validator.Struct(data)
}

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
