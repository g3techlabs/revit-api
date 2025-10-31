package validation

import (
	"mime/multipart"
	"net/mail"
	"reflect"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	upperCase = regexp.MustCompile(`[A-Z]`)
	lowerCase = regexp.MustCompile(`[a-z]`)
	minLength = regexp.MustCompile(`.{8,}`)
)

func Password(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	return passwordRegexCheck(password)
}

func passwordRegexCheck(s string) bool {
	return upperCase.MatchString(s) &&
		lowerCase.MatchString(s) &&
		minLength.MatchString(s)
}

func ProfilePic(fl validator.FieldLevel) bool {
	file, ok := fl.Field().Interface().(*multipart.FileHeader)
	if !ok || file == nil {
		return false
	}

	contentType := file.Header.Get("Content-Type")
	switch contentType {
	case "image/png", "image/jpeg", "image/webp":
		return true
	default:
		return false
	}
}

func IdentifierTypeValidation(sl validator.StructLevel) {
	val := reflect.ValueOf(sl.Current().Interface())

	identifierField := val.FieldByName("Identifier")
	identifierTypeField := val.FieldByName("IdentifierType")

	if !identifierField.IsValid() || !identifierTypeField.IsValid() {
		return
	}

	identifier := identifierField.String()
	identifierType := identifierTypeField.String()

	_, err := mail.ParseAddress(identifier)

	switch identifierType {
	case "email":
		if err != nil {
			sl.ReportError(identifier, "Identifier", "identifier", "notanemail", "")
		}

	case "nickname":
		if err == nil {
			sl.ReportError(identifier, "Identifier", "identifier", "notanickname", "")
			return
		}
		if len(identifier) < 3 {
			return
		}

	default:
		return
	}
}
