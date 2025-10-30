package errors

import "github.com/g3techlabs/revit-api/response"

func InvalidDateValue() error {
	return &response.CustomError{
		Message:    "The date value must be at least 15 minutes from now",
		StatusCode: 400,
	}
}
