package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	upperCase = regexp.MustCompile(`[A-Z]`)
	lowerCase = regexp.MustCompile(`[a-z]`)
	minLength = regexp.MustCompile(`.{8,}`)
)

func password(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	return passwordRegexCheck(password)
}

func passwordRegexCheck(s string) bool {
	return upperCase.MatchString(s) &&
		lowerCase.MatchString(s) &&
		minLength.MatchString(s)
}
