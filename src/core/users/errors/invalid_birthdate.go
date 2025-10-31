package errors

import "github.com/g3techlabs/revit-api/src/response"

func InvalidBirthdateFormat() *response.CustomError {
	return &response.CustomError{
		StatusCode: 400,
		Message:    "Invalid format date",
	}
}
