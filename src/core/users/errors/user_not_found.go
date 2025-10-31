package errors

import "github.com/g3techlabs/revit-api/src/response"

func UserNotFound(message string) error {
	return &response.CustomError{
		Message:    message,
		StatusCode: 404,
	}
}
