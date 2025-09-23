package utils

import (
	"unicode"
)

func HasUpperCase(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}

func HasLowerCase(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) {
			return true
		}
	}
	return false
}

func CollectErrors(errs ...error) []string {
	var out []string
	for _, err := range errs {
		if err != nil {
			out = append(out, err.Error())
		}
	}
	return out
}
