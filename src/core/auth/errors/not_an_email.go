package errors

import "github.com/g3techlabs/revit-api/src/response"

func NotAnEmail() error {
	return &response.CustomError{
		Message:    "Identifier is not an email",
		StatusCode: 400,
	}
}
