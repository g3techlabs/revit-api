package errors

import "github.com/g3techlabs/revit-api/src/response"

func InvalidDateFormat() error {
	return &response.CustomError{
		Message:    "Invalid date format",
		StatusCode: 400,
	}
}
